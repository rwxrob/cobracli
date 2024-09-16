package cobracli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func countLeadingSpaces(s string) int {
	count := 0
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			break
		}
		count++
	}
	return count
}

// Includes initialization of slog
// Includes sample of embedded docs using embed package
// Assumes run from a git repo
func InitCommands(a io.Reader) error {
	var err error
	scanner := bufio.NewScanner(a)
	path := []string{`cmd`}
	var lspaces int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			spaces := countLeadingSpaces(line)
			command := strings.TrimSpace(line)
			switch {
			case spaces > lspaces:
				path = append(path, command)
			case spaces < lspaces:
				path = path[:len(path)-1]
			case spaces == lspaces:
				path[len(path)-1] = command
			}
			lspaces = spaces
		}
		fmt.Println(filepath.Join(path...))
		if err = os.MkdirAll(filepath.Join(path...), 0755); err != nil {
			return err
		}
		// TODO
	}
	return err
}

func InitInternal() error {
	// TODO
	return nil
}

func InitModule() error {
	// TODO run go mod init BLAH
	return nil
}

func RemindGoWork() error {
	// TODO
	return nil
}
