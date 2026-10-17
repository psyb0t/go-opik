// Command patch-spec normalizes YAML scalar tags for strict OpenAPI loaders.
package main

import (
	"os"
	"strings"

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

	fixDescriptions(node)

	for _, child := range node.Content {
		fixValueTags(child)
	}
}

func fixDescriptions(node *yaml.Node) {
	if node.Kind != yaml.MappingNode {
		return
	}

	const mappingPairSize = 2

	for index := 0; index < len(node.Content); index += mappingPairSize {
		key, value := node.Content[index], node.Content[index+1]
		if key.Value != "description" && key.Value != "summary" {
			continue
		}

		if value.Kind != yaml.ScalarNode {
			continue
		}

		// Keep generated comments in the repository's plain punctuation style.
		value.Value = strings.NewReplacer(
			" \u2014 ", ", ",
			" \u2013 ", ", ",
			" - ", ", ",
		).Replace(value.Value)
	}
}
