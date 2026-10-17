package opik_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-opik"
)

// This is the README usage example. It requires cloud credentials.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := opik.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", os.Getenv("OPIK_API_KEY"))
			req.Header.Set("Comet-Workspace", os.Getenv("OPIK_WORKSPACE"))

			return nil
		},
	)

	client, err := opik.NewClientWithResponses(
		"https://www.comet.com/opik/api",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.FindProjectsWithResponse(
		context.Background(),
		&opik.FindProjectsParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
