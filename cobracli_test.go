package cobracli_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/cobracli"
)

func ExampleInitCommands() {

	buf := strings.NewReader(`
cmdname
  treetop
    branch
    descendent
      waydown
  bar
    justonebranch
  nosubs
`)

	if err := cobracli.InitCommands(buf); err != nil {
		fmt.Println(err)
		return
	}

	// Output:
	// foo
}
