package testmod

import "testing"

func TestHello(t *testing.T) {
    want := "Привет, Ivan!"

    if got, _ := Hello("Ivan", "ru"); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
