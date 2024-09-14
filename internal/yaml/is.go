package yaml

import (
	"gopkg.in/yaml.v3"
)

func IsDocument(n *yaml.Node) bool   { return n != nil && n.Kind == yaml.DocumentNode }
func IsSequence(n *yaml.Node) bool   { return n != nil && n.Kind == yaml.SequenceNode }
func IsMapping(n *yaml.Node) bool    { return n != nil && n.Kind == yaml.MappingNode }
func IsScalar(n *yaml.Node) bool     { return n != nil && n.Kind == yaml.ScalarNode }
func IsAlias(n *yaml.Node) bool      { return n.Alias != nil && n.Kind == yaml.AliasNode }
func IsCollection(n *yaml.Node) bool { return IsMapping(n) || IsSequence(n) }
func IsNull(n *yaml.Node) bool       { return n != nil && n.Tag == `!!null` }
func IsString(n *yaml.Node) bool     { return n != nil && n.Tag == `!!str` }
func IsInt(n *yaml.Node) bool        { return n != nil && n.Tag == `!!int` }
func IsFloat(n *yaml.Node) bool      { return n != nil && n.Tag == `!!float` }
func IsNumber(n *yaml.Node) bool     { return IsInt(n) || IsFloat(n) }
func IsBool(n *yaml.Node) bool       { return n != nil && n.Tag == `!!bool` }
func IsMerge(n *yaml.Node) bool      { return n != nil && n.Tag == `!!merge` }
func IsBinary(n *yaml.Node) bool     { return n != nil && n.Tag == `!!binary` }
func IsTimeStamp(n *yaml.Node) bool  { return n != nil && n.Tag == `!!timestamp` }

func IsMapKey(node, from *yaml.Node, contentIndex int) bool {
	return node != nil && from != nil && contentIndex%2 == 0 && from.Kind == yaml.MappingNode
}

func IsMapVal(node, from *yaml.Node, contentIndex int) bool {
	return node != nil && from != nil && contentIndex%2 != 0 && from.Kind == yaml.MappingNode
}

// MapKeyFor returns from.Content[contentIndex-1].Value and should only be
// passed a valnode that passes [IsMapVal]. Returns empty string if
// unable to lookup.
func MapKeyFor(valnode, from *yaml.Node, contentIndex int) string {
	if valnode == nil || from == nil || from.Content == nil {
		return ""
	}
	return from.Content[contentIndex-1].Value
}

// MapValFor returns from.Content[contentIndex+1].Value and should only be
// called passed a keynode that passes [IsMapKey]. Returns empty string if
// unable to lookup.
func MapValFor(keynode, from *yaml.Node, contentIndex int) string {
	if keynode == nil || from == nil || from.Content == nil {
		return ""
	}
	return from.Content[contentIndex-1].Value
}
