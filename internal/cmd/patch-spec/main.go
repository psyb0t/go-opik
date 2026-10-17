// Command patch-spec normalizes YAML scalar tags for strict OpenAPI loaders.
package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	sourceSpec  = "api/openapi.yaml"
	patchedSpec = "api/openapi.fixed.yaml"
	specMode    = 0o644
	valueTag    = "tag:yaml.org,2002:value"
)

func main() {
	data, err := os.ReadFile(sourceSpec)
	if err != nil {
		panic(err)
	}

	var document yaml.Node
	if err := yaml.Unmarshal(data, &document); err != nil {
		panic(err)
	}

	fixValueTags(&document)

	data, err = yaml.Marshal(&document)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(patchedSpec, data, specMode); err != nil {
		panic(err)
	}
}

func fixValueTags(node *yaml.Node) {
	// YAML 1.1 resolves bare = as !!value, although these enum entries
	// describe literal strings. Quoting preserves the API wire value.
	if node.Kind == yaml.ScalarNode &&
		(node.Tag == valueTag || node.Value == "=") {
		node.Tag = "!!str"
		node.Style = yaml.DoubleQuotedStyle
	}

	for _, child := range node.Content {
		fixValueTags(child)
	}
}
