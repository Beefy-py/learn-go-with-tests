package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeats a string 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	result := Repeat("b", 3)
	println(result) // Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
