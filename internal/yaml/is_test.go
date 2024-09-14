package yaml_test

import (
	"fmt"

	"github.com/rwxrob/cobracli/internal/yaml"
	yamlv3 "gopkg.in/yaml.v3"
)

func ExampleIsBool() {
	values := []any{
		true,
		false,
		`true`,
		`True`,
		`Yes`,
		`yes`,
	}
	for _, val := range values {
		node, _ := yaml.ToNode(val)
		fmt.Println(yaml.IsBool(node))
	}
	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleIsBool_from_Node() {
	node := new(yamlv3.Node)
	node.Kind = yamlv3.ScalarNode
	node.Tag = `!!bool`
	node.Value = `True`
	fmt.Println(yaml.IsBool(node))
	fmt.Print(yaml.ToString(node))
	// Output:
	// true
	// True
}

func ExampleIsBool_from_Yes() {
	node := new(yamlv3.Node)
	node.Kind = yamlv3.ScalarNode
	node.Tag = `!!bool`
	node.Value = `Yes`
	fmt.Println(yaml.IsBool(node))
	fmt.Print(yaml.ToString(node))
	// Output:
	// true
	// !!bool Yes
}
