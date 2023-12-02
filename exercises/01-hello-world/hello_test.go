package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test 1: say hello to a person", func(t *testing.T) {
		got := Hello("Adam", "")
		want := "Hello, Adam!!!"
		assertCorrectMessage(t, got, want)
		//		if got != want {
		//			t.Errorf("got %q want %q", got, want)
		//		}
	})
	t.Run("Test 2: say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!!!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 3: say hello to a person in Klingon", func(t *testing.T) {
		got := Hello("Worf", "Klingon")
		want := "nuqneH, Worf!!!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test 4: say hello to a person in Romulan", func(t *testing.T) {
		got := Hello("Worf", "Romulan")
		want := "jolan'tru, Worf!!!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
