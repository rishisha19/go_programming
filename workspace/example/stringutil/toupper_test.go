package stringutil

import "testing"

func TestToUpper(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "HELLO, WORLD"},
		{"sunny DAy", "SUNNY DAY"},
		{"", ""},
	} {
		got := ToUpper(c.in)
		if got != c.want {
			t.Errorf("ToUpper(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
