package random

import "testing"

func TestRandomString(t *testing.T) {
	t.Run("Test random string", func(t *testing.T) {
		got := NewRandomString(10)
		if len(got) != 10 {
			t.Errorf("NewRandomString() = %v, want %v", len(got), 10)
		}
	})
}
