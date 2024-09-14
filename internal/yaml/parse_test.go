package yaml_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/cobracli/internal/yaml"
)

func ExampleParse() {

	r := strings.NewReader(`
%YAML 1.1
---
title: YAML Ain't Markup Language
foo: bar
%YAML 1.1
---
another: thing
one: 1
%YAML 1.1
%TAG ! tag:clarkevans.com,2002:
---
some: random thing
slice:
  # comment
  # about
  - one
  - two
--- I'm a valid scalar string
---
- one
- two
- three
--- !!str
iam: broken
`)

	docs, err := yaml.Parse(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, doc := range docs {

		fmt.Printf("Kind: %q ShortTag: %q LongTag: %q Value: %q\n",
			yaml.Kinds[doc.Content[0].Kind],
			doc.Content[0].ShortTag(),
			doc.Content[0].LongTag(),
			doc.Content[0].Value,
		)
	}

	// Note that the last item is wrong in the yaml.v3 implementation and
	// has been for years.

	// Output:
	// Kind: "Mapping" ShortTag: "!!map" LongTag: "tag:yaml.org,2002:map" Value: ""
	// Kind: "Mapping" ShortTag: "!!map" LongTag: "tag:yaml.org,2002:map" Value: ""
	// Kind: "Mapping" ShortTag: "!!map" LongTag: "tag:yaml.org,2002:map" Value: ""
	// Kind: "Scalar" ShortTag: "!!str" LongTag: "tag:yaml.org,2002:str" Value: "I'm a valid scalar string"
	// Kind: "Sequence" ShortTag: "!!seq" LongTag: "tag:yaml.org,2002:seq" Value: ""
	// Kind: "Mapping" ShortTag: "!!str" LongTag: "tag:yaml.org,2002:str" Value: ""

}

func ExampleParseString() {
	docs, err := yaml.ParseString(`--- something`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yaml.Kinds[docs[0].Content[0].Kind])
	// Output:
	// Scalar
}

func ExampleParseString_check_Docs() {
	docs, err := yaml.ParseString("")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(docs) == 0 {
		fmt.Println(`Empty.`)
		return
	}

	// Output:
	// Empty.
}

func ExampleParseBytes() {
	docs, err := yaml.ParseBytes([]byte(`--- something`))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yaml.Kinds[docs[0].Content[0].Kind])
	// Output:
	// Scalar
}
