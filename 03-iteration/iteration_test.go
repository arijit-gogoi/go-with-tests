package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat default times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"

		assertEqual(t, repeated, expected)
	})

	t.Run("repeat 2 times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		assertEqual(t, repeated, expected)
	})

	t.Run("repeat negative times", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "a"

		assertEqual(t, repeated, expected)
	})
}

func assertEqual(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected: %q\nrepeated: %q", expected, repeated)
	}
}

func ExampleRepeat() {
	out := Repeat("b", -1)
	fmt.Println(out)
	// Output: b
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
