package opik_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/psyb0t/go-opik"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey          = "EXAMPLE-DO-NOT-USE"
	testWorkspace       = "test-workspace"
	testProjectName     = "test-project"
	testTraceID         = "00000000-0000-4000-8000-000000000001"
	projectsPath        = "/api/v1/private/projects"
	jsonContentType     = "application/json"
	authorizationHeader = "Authorization"
	workspaceHeader     = "Comet-Workspace"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *opik.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := opik.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(authorizationHeader, testAPIKey)
			req.Header.Set(workspaceHeader, testWorkspace)

			return nil
		},
	)

	client, err := opik.NewClientWithResponses(srv.URL+"/api", auth)
	require.NoError(t, err)

	return client
}

func assertAuth(t *testing.T, r *http.Request) {
	t.Helper()
	assert.Equal(t, testAPIKey, r.Header.Get(authorizationHeader))
	assert.Equal(t, testWorkspace, r.Header.Get(workspaceHeader))
}

func writeJSON(t *testing.T, w http.ResponseWriter, body string) {
	t.Helper()
	w.Header().Set("Content-Type", jsonContentType)
	_, err := w.Write([]byte(body))
	assert.NoError(t, err)
}

func TestFindProjectsPagination(t *testing.T) {
	t.Parallel()

	const pageSize = int32(2)

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, projectsPath, r.URL.Path)
		assert.Equal(t, "2", r.URL.Query().Get("size"))
		assert.Equal(t, "test & project", r.URL.Query().Get("name"))
		assertAuth(t, r)

		switch r.URL.Query().Get("page") {
		case "1":
			writeJSON(t, w, `{
				"page":1,"size":2,"total":3,
				"content":[{"name":"first"},{"name":"second"}]
			}`)
		case "2":
			writeJSON(t, w, `{
				"page":2,"size":2,"total":3,"content":[{"name":"third"}]
			}`)
		default:
			writeJSON(t, w, `{"page":3,"size":2,"total":3,"content":[]}`)
		}
	})

	var names []string

	for page := int32(1); page <= 3; page++ {
		res, err := client.FindProjectsWithResponse(
			context.Background(),
			&opik.FindProjectsParams{
				Page: &page,
				Size: new(pageSize),
				Name: new("test & project"),
			},
		)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode())
		require.NotNil(t, res.JSON200)
		require.NotNil(t, res.JSON200.Content)
		assert.Equal(t, int64(3), *res.JSON200.Total)
		assert.Equal(t, page, *res.JSON200.Page)

		for _, project := range *res.JSON200.Content {
			names = append(names, project.Name)
		}
	}

	assert.Equal(t, []string{"first", "second", "third"}, names)
}

func TestGetTrace(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/v1/private/traces/"+testTraceID, r.URL.Path)
		assert.Equal(t, "true", r.URL.Query().Get("strip_attachments"))
		assertAuth(t, r)
		writeJSON(t, w, `{
			"id":"00000000-0000-4000-8000-000000000001",
			"name":"test-trace","start_time":"2026-10-17T00:00:00Z",
			"total_estimated_cost":0.123456789123,"end_time":null
		}`)
	})

	res, err := client.GetTraceByIdWithResponse(
		context.Background(),
		uuid.MustParse(testTraceID),
		&opik.GetTraceByIdParams{StripAttachments: new(true)},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	assert.Equal(t, "test-trace", *res.JSON200.Name)
	assert.Equal(t, uuid.MustParse(testTraceID), *res.JSON200.Id)
	assert.Equal(t, 0.123456789123, *res.JSON200.TotalEstimatedCost)
	assert.Nil(t, res.JSON200.EndTime)
}

func TestCreateProject(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, projectsPath, r.URL.Path)
		assert.Equal(t, jsonContentType, r.Header.Get("Content-Type"))
		assertAuth(t, r)

		var body opik.CreateProjectJSONRequestBody
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, testProjectName, body.Name)
		w.Header().Set("Location", projectsPath+"/"+testTraceID)
		w.WriteHeader(http.StatusCreated)
	})

	res, err := client.CreateProjectWithResponse(
		context.Background(),
		opik.CreateProjectJSONRequestBody{Name: testProjectName},
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode())
	assert.Equal(t, projectsPath+"/"+testTraceID,
		res.HTTPResponse.Header.Get("Location"))
}

func TestUnauthorized(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		key  string
	}{
		{name: "missing"},
		{name: "malformed", key: "invalid"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := newTestClient(t,
				func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, tc.key, r.Header.Get(authorizationHeader))
					w.WriteHeader(http.StatusUnauthorized)
				},
			)
			res, err := client.FindProjectsWithResponse(
				context.Background(), &opik.FindProjectsParams{},
				func(_ context.Context, req *http.Request) error {
					req.Header.Set(authorizationHeader, tc.key)

					return nil
				},
			)
			require.NoError(t, err)
			assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
			assert.Nil(t, res.JSON200)
		})
	}
}

func TestInvalidJSON(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, `{`)
	})
	_, err := client.FindProjectsWithResponse(
		context.Background(), &opik.FindProjectsParams{},
	)

	var syntaxError *json.SyntaxError
	require.ErrorAs(t, err, &syntaxError)
}

func TestCanceledContext(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(_ http.ResponseWriter, _ *http.Request) {
		t.Error("canceled request reached the server")
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.FindProjectsWithResponse(ctx, &opik.FindProjectsParams{})
	require.ErrorIs(t, err, context.Canceled)
}

func TestEqualityOperator(t *testing.T) {
	t.Parallel()

	var condition opik.ScoreCondition

	err := json.Unmarshal(
		[]byte(`{"operator":"=","score_name":"quality","value":0.5}`),
		&condition,
	)
	require.NoError(t, err)
	assert.Equal(t, "=", string(condition.Operator))
	assert.Equal(t, "0.5", strconv.FormatFloat(condition.Value, 'f', -1, 64))
}
