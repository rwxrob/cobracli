package internal_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/cobracli/internal"
)

func ExampleCountLeadingSpaces() {
	fmt.Println(internal.CountLeadingSpaces(`  some`))
	fmt.Println(internal.CountLeadingSpaces(`some`))
	fmt.Println(internal.CountLeadingSpaces(`          some`))

	// Output:
	// 2
	// 0
	// 10
}

func ExampleIndentedToSlices() {

	in := strings.NewReader(`
usr
  bin
  local
    bin
var
  local
	log
	game
	  gui
		tui
`)

	paths := internal.IndentedToSlices(in)
	for _, path := range paths {
		fmt.Println(path)
	}

	// Output:
	// [usr]
	// [usr bin]
	// [usr local]
	// [usr local bin]
	// [var]
	// [var local]
	// [var log]
	// [var game]
	// [var game gui]
	// [var game tui]
}
