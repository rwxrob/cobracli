package yaml

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

var Kinds = map[yaml.Kind]string{
	yaml.DocumentNode: `Document`,
	yaml.SequenceNode: `Sequence`,
	yaml.MappingNode:  `Mapping`,
	yaml.ScalarNode:   `Scalar`,
	yaml.AliasNode:    `Alias`,
}

// ToString is a convenience function for [yaml.Marshal] that casts the
// byte slice into a string and logs the error instead of returning it.
func ToString(it any) string {
	byt, err := yaml.Marshal(it)
	if err != nil {
		log.Println(err)
	}
	return string(byt)
}

// Print prints the marshaled bytes as a string to os.Stdout. If any
// error occurs the "null" string (no quotes) is printed followed by
// a line return.
func Print(it any) {
	byt, err := yaml.Marshal(it)
	if err != nil {
		fmt.Println(`null`)
	}
	fmt.Print(string(byt))
}

// ToDocNode returns a reference to a [gopkg.in/yaml.v3.Node] by first
// marshaling the argument into a YAML document string and then
// unmarshaling into a new Node and returning its reference along with
// the containing DocumentNode. When just the Node itself is wanted use
// [ToNode] instead.
func ToDocNode(it any) (*yaml.Node, error) {
	byt, err := yaml.Marshal(it)
	if err != nil {
		return nil, err
	}
	node := new(yaml.Node)
	if err := yaml.Unmarshal(byt, node); err != nil {
		return nil, err
	}
	return node, nil
}

// ToNode returns a reference to a [gopkg.in/yaml.v3.Node] by first
// marshaling the argument into a YAML document string and then
// unmarshaling into a new Node and returning its reference (Content[0]).
// The Node can then be transformed or queried by a [WalkFunc] with
// [Walk]. When the encapsulating document is also wanted use [ToDocNode]
// instead.
func ToNode(it any) (*yaml.Node, error) {
	node, err := ToDocNode(it)
	if err != nil {
		return nil, err
	}
	return node.Content[0], nil
}
