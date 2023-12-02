package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d' instead", expected, sum)
	}
}

func ExampleAdder() {
	sum := Add(10, 20)
	fmt.Println(sum)
	// Output: 30
}
