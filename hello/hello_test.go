package main

import "testing"


// write a test function here for each function you have written
func TestHello(t *testing.T) {
	// declaring some vars:
	got := Hello("Chris")
	want := "Hello, Chris"

	// if statements and string formatting are pretty standard:
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

