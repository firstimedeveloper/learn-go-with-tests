package main

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Run("updating existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		newDef := "new definition for the word"

		err := dictionary.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDef)
	})
	//a use case of update where it should fail, as the word doesn't exist
	t.Run("updating a word that does not exist (fails)", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDef := "new definition for the word"

		err := dictionary.Update(word, newDef)
		assertError(t, err, ErrWordNotExists)
		//assertDefinition(t, dictionary, word, newDef)
	})
}

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

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	//t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if def != got {
		t.Errorf("got '%s' want '%s'", got, def)
	}
}

func assertError(t *testing.T, got, want error) {
	//t.Helper()
	if got != want {
		t.Errorf("got error '%s' want error '%s'", got, want)
	}
}
