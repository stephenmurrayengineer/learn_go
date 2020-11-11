package main

import "testing"


// write a test function here for each function you have written
func TestHello(t *testing.T) {

	// refactor tests into an assertion to remove duplication
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// helper func displays line func is called:
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// test case where there is an arg to hello function:
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	// test case where there is no arg
	t.Run("say hello world, when func takes an empty string", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// test case for spanish
	t.Run("say hello world in spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world in french", func(t *testing.T){
		got := Hello("Jaques", "French")
		want := "Bonjour, Jaques"
		assertCorrectMessage(t, got, want)
	})

}

