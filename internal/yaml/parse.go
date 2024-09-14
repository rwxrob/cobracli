package yaml

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse parses a multidoc YAML stream into a slice of rooted [yaml.Node]
// references, each being of type [yaml.DocumentNode]. Parsing stops if
// any error is encountered during decoding but all nodes that were
// successfully parsed up to that point are still returned with the
// error. A nil argument causes a panic. Attempting to read an empty
// document (only white space, for example) does not produce an error
// but does result in an empty slice so be sure to check its length
// before using.
func Parse(r io.Reader) ([]*yaml.Node, error) {
	nodes := []*yaml.Node{}
	d := yaml.NewDecoder(r)
	for {
		node := new(yaml.Node)
		err := d.Decode(node)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// ParseString calls [Parse] after creating an io.Reader for it.
func ParseString(a string) ([]*yaml.Node, error) {
	return Parse(strings.NewReader(a))
}

// ParseBytes calls [Parse] after creating an io.Reader for it.
func ParseBytes(a []byte) ([]*yaml.Node, error) {
	return Parse(bytes.NewReader(a))
}
