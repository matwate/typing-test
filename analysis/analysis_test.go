package analysis

import (
	"testing"
)

func TestDistance(t *testing.T) {
	t.Run("Distance", func(t *testing.T) {
		got := Distance("kitten", "sitting")
		want := 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestLoadAllowedWords(t *testing.T) {
	t.Run("LoadAllowedWords", func(t *testing.T) {
		// We'll just test if the function runs without error
		load_allowed_words()
	})
}

func TestFindProperTarget(t *testing.T) {
	t.Run("FindProperTarget", func(t *testing.T) {
		load_allowed_words()
		got := find_proper_target("kitten")
		want := "kitten"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got = find_proper_target("kwtten")
		want = "kitten"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
