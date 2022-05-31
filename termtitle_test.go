package termtitle

import (
	"testing"
)

func TestSetAndGetTitle(t *testing.T) {
	titles := []string{
		"",
		"FOO",
		"Foo",
		"1234567890",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"[]{}Foo",
		"\u0100\u00FF\u00F1\u00E3",
		"\u016E\u0174\u0197\u01AA\u01BA",
		"",
	}

	for _, title := range titles {
		if err := SetTitle(title); err != nil {
			t.Errorf("SetTitle(%q) failed: %v", title, err)
		}

		got, err := GetTitle()

		if err != nil {
			t.Errorf("GetTitle() failed: %s", err)
		}

		if got != title {
			t.Errorf("got title %q, want %q", got, title)
		}
	}

}
