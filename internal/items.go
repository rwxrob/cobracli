package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"slices"
	"strings"
)

type LineItem struct {
	Raw    string
	Spaces int
	Text   string
	Parent *LineItem
}

func (t LineItem) String() string {
	if t.Parent != nil {
		return fmt.Sprintf("%v %v.%v", t.Spaces, t.Parent.Text, t.Text)
	}
	return fmt.Sprintf("%v %v", t.Spaces, t.Text)
}

func (t *LineItem) IsLowerThan(a *LineItem) bool   { return t.Spaces > a.Spaces }
func (t *LineItem) IsHigherThan(a *LineItem) bool  { return a.Spaces > t.Spaces }
func (t *LineItem) IsSameLevelAs(a *LineItem) bool { return a.Spaces == t.Spaces }

func (t *LineItem) Root() *LineItem {
	var cur *LineItem
	for cur = t; cur.Parent != nil; cur = cur.Parent {
	}
	return cur
}

func (t *LineItem) AsLineItems() []*LineItem {
	items := []*LineItem{}
	cur := t
	for cur = t; cur.Parent != nil; cur = cur.Parent {
		items = append(items, cur)
	}
	items = append(items, cur)
	slices.Reverse(items)
	return items
}

func (t *LineItem) AsSlice() []string {
	texts := []string{}
	for _, item := range t.AsLineItems() {
		texts = append(texts, item.Text)
	}
	return texts
}

func (t *LineItem) AsPath() string   { return filepath.Join(t.AsSlice()...) }
func (t *LineItem) AsDotted() string { return strings.Join(t.AsSlice(), ".") }

func IndentedToItems(in io.Reader) []*LineItem {
	items := []*LineItem{}

	var prev *LineItem
	last := map[int]*LineItem{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {

		// parse a line into a line item (without empty fields for now)
		item := new(LineItem)
		item.Raw = scanner.Text()
		item.Text = strings.TrimSpace(item.Raw)
		if len(item.Text) == 0 {
			continue
		}
		item.Spaces = CountLeadingSpaces(item.Raw)

		// infer the Fields
		switch {
		case prev == nil:
			// ignore

		case item.IsLowerThan(prev):
			item.Parent = prev

		case item.IsSameLevelAs(prev):
			item.Parent = prev.Parent

		case item.IsHigherThan(prev):
			if lastpeer, has := last[item.Spaces]; has {
				item.Parent = lastpeer.Parent
			} else {
				log.Print(`syntax error, unmatched level`)
				return items
			}

		}

		last[item.Spaces] = item
		prev = item
		items = append(items, item)
	}

	return items
}
