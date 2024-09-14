package yaml_test

import (
	"fmt"

	"github.com/rwxrob/cobracli/internal/yaml"
	yamlv3 "gopkg.in/yaml.v3"
)

func ExampleToString() {
	types := []any{"foo", 1.3, -390, 42, []string{"one", "two"}}
	for _, arg := range types {
		fmt.Print(yaml.ToString(arg))
	}
	// Output:
	// foo
	// 1.3
	// -390
	// 42
	// - one
	// - two
}

func ExamplePrint() {
	types := []any{
		"foo",
		1.3, -390, 42,

		[]string{"one", "two"}, // Sequence of same scalars
		[]any{"one", 2, 3.1},   // Sequence of different

		map[int]string{1: "one", 2: "two"}, // Mapping

		struct {
			Num int
			Str string
		}{12, "twelve"},

		nil,
	}
	for _, arg := range types {
		yaml.Print(arg)
	}
	// Output:
	// foo
	// 1.3
	// -390
	// 42
	// - one
	// - two
	// - one
	// - 2
	// - 3.1
	// 1: one
	// 2: two
	// num: 12
	// str: twelve
	// null

}

func ExampleToDocNode() {
	node, err := yaml.ToDocNode("foo")
	if err != nil {
		fmt.Println(err)
	}

	// gets marshalled as same string as input
	yaml.Print(node)

	// content always encapsulated in a DocumentNode Kind
	fmt.Println(node.Kind == yamlv3.DocumentNode)

	// first in Content slice has value
	fmt.Println(node.Content[0].Kind == yamlv3.ScalarNode)

	// synonymous
	fmt.Println(node.Content[0].Tag)
	fmt.Println(node.Content[0].ShortTag())

	// long form as specified by YAML spec
	fmt.Println(node.Content[0].LongTag())

	// Output:
	// foo
	// true
	// true
	// !!str
	// !!str
	// tag:yaml.org,2002:str

}

func ExampleToNode_bool_Tag_Still_String() {
	node, _ := yaml.ToNode(`!!bool Yes`)
	fmt.Println(node.Tag)
	fmt.Println(node.Value)
	// Output:
	// !!str
	// !!bool Yes
}
