package testmod

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, Ivan!"

    if got := Hello("Ivan"); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
