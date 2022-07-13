package main

import (
	"testing"
	hello "github.com/tang9ian/gomodule-demo"
)

func TestHello(t *testing.T) {
	name := "vim-go!"
	got := hello.Hello(name)
	want := "Hello, vim-go!"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
