package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleAdd() {
	result := Add(1, 2)
	fmt.Println(result)
	// Output: 3
}
