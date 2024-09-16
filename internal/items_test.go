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

func ExampleLineItem_Root() {

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

	items := internal.IndentedToItems(in)

	fmt.Println(items[3].Root().Text)
	fmt.Println(items[8].Root().Text)

	fmt.Println()

	// Output:
	// usr
	// var
}

func ExampleLineItem_AsPath() {

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

	items := internal.IndentedToItems(in)

	fmt.Println(items[3].AsPath())
	fmt.Println(items[8].AsPath())

	fmt.Println()

	// Output:
	// usr/local/bin
	// var/game/gui
}

func ExampleLineItem_AsDotted() {

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

	items := internal.IndentedToItems(in)

	fmt.Println(items[3].AsDotted())
	fmt.Println(items[8].AsDotted())

	fmt.Println()

	// Output:
	// usr.local.bin
	// var.game.gui
}

func ExampleIndentedToItems_String() {

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

	items := internal.IndentedToItems(in)
	for _, item := range items {
		fmt.Println(item)
	}

	// Output:
	// 0 usr
	// 2 usr.bin
	// 2 usr.local
	// 4 local.bin
	// 0 var
	// 2 var.local
	// 2 var.log
	// 2 var.game
	// 4 game.gui
	// 4 game.tui

}

/*
func ExampleIndentedToItems() {

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

	items := internal.IndentedToItems(in)
	for _, item := range *items {
		fmt.Println(item.Path())
	}

	// Output:
	// usr
	// usr/bin
	// usr/local
	// usr/local/bin
	// var
	// var/local
	// var/log
	// var/game
	// var/game/gui
	// var/game/tui
}
*/
