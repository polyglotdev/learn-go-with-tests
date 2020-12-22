package main

import "testing"

func Test_main(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
