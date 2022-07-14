package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, vim-go!"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
