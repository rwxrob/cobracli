package yaml

import (
	"gopkg.in/yaml.v3"
)

// WalkFunc is any function passed to [Walk] or [WalkFrom] that is
// applied to each [gopkg.in/yaml.v3.Node] (in no particular order).
// Such functions can be used as validators and aggregators by enclosing
// variables within them.
// and the information about
// how the Node is contained and used sufficient to determine its
// specified YAML type (see is.go). The containing (from) Node may be
// nil for root Nodes. The contentIndex must be -1 if the Content field
// is undefined.
type WalkFunc func(node, from *yaml.Node, contentIndex int) error

// Walk traverses every node of the root [yaml.Node] passed. Currently,
// this is done in a modified depth-first way that can produce output
// somewhat consistent with the YAML representation but the method of
// traversal should never be relied upon as it may change. The only
// guarantee is that the [WalkFunc] is applied to each and every node
// exactly once.  Walk is exactly equivalent to WalkFrom(root,nil,-1,fn).
func Walk(root *yaml.Node, fn WalkFunc) error {
	return WalkFrom(root, nil, -1, fn)
}

// WalkFrom is like [Walk] but passes the information necessary to work
// with the containing (from) Node and contentIndex providing essential
// context for determining the type of node.
func WalkFrom(node, from *yaml.Node, contentIndex int, fn WalkFunc) error {
	if err := fn(node, from, contentIndex); err != nil {
		return err
	}
	switch node.Kind {
	case yaml.DocumentNode:
		if err := WalkFrom(node.Content[0], node, 0, fn); err != nil {
			return err
		}
	case yaml.MappingNode, yaml.SequenceNode:
		for i, n := range node.Content {
			if err := WalkFrom(n, node, i, fn); err != nil {
				return err
			}
		}
	case yaml.AliasNode:
		if err := WalkFrom(node.Alias, node, -1, fn); err != nil {
			return err
		}
	}
	return nil
}
