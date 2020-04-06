package sdk

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestInitialize(t *testing.T) {

	if got := Initialize("", ""); got != nil {
		t.Errorf("Error Initializing() = %s", got)
	}
}
