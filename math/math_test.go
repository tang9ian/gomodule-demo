package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(5, 5)
	want := 10
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	// t.Fatal("not implemented")
}
