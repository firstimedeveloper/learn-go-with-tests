package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		def := "this is just a test"
		err := dictionary.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, def)
	})
	t.Run("atempt to add an existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, def)

	})
}

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	//t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if def != got {
		t.Errorf("got %s want %s", got, def)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	// t.Run("known word", func(t *testing.T) {
	// 	got, _ := dictionary.Search("test")
	// 	want := "this is just a test"

	// 	assertError(t, got, want)
	// })

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected an error")
		}

		assertError(t, got, ErrNotFound)
	})

}

func assertError(t *testing.T, got, want error) {
	//t.Helper()
	if got != want {
		t.Errorf("got error '%s' want error '%s'", got, want)
	}
}
