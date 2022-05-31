//go:build !windows

package termtitle

import "fmt"

func setTitle(title string) error {
	_, err := fmt.Print("\033]0;" + title + "\007")
	return err
}

func getTitle() (string, error) {
	panic("termtitle: GetTitle not implemented on this platform")
}
