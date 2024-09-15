package internal

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func CountLeadingSpaces(s string) int {
	count := 0
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			break
		}
		count++
	}
	return count
}

func IndentedToSlices(in io.Reader) (paths [][]string) {
	scanner := bufio.NewScanner(in)
	var lspaces int
	current := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		this := strings.TrimSpace(line)
		if len(this) == 0 {
			continue
		}
		// FIXME redo all of this with spaces as indentation level indicator
		spaces := CountLeadingSpaces(line)
		switch {
		case spaces == lspaces: // peer
			newpath := []string{}
			if len(current) == 0 {
				newpath = []string{this}
			} else {
				// required to force a new slice else bork references
				newpath = append(newpath, current[:len(current)-1]...)
				newpath = append(newpath, this)
			}
			paths = append(paths, newpath)
			current = newpath
		case spaces > lspaces: // new level
			newpath := []string{}
			newpath = append(newpath, current...)
			newpath = append(newpath, this)
			current = newpath
			paths = append(paths, newpath)
		case spaces < lspaces: // back
			newpath := []string{}
			// FIXME might be more than just one level back
			newpath = append(newpath, current[:len(current)-1]...)
			current = newpath
			paths = append(paths, newpath)
		}
		lspaces = spaces
	}
	return
}
