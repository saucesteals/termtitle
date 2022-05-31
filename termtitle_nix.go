//go:build !windows

package termtitle

import "fmt"

func setTitle(title string) error {
	fmt.Print("\033]0;" + title + "\007")
}

func getTitle() (string, error) {
	panic("termtitle: GetTitle not implemented on this platform")
}
