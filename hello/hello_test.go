package hello

import (
	"fmt"
	"testing"
)

func TestHelloWithoutName(t *testing.T) {
	want := "Hello, "
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	alex.Name = "Alex"
	want := fmt.Sprintf("Hello, %s", alex.Name)
	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}
