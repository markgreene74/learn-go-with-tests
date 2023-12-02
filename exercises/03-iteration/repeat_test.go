package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Test 1: repeat with no extra args", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 1: repeat with extra argument 1", func(t *testing.T) {
		got := Repeat("z", 1)
		want := "z"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 1: repeat with extra argument 7", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("z", 3)
	fmt.Println(repeated)
	// Output: zzz
}
