package search

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		want := "this is just a test"

		assertDefinition(t, dict, "test", want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		_, err := dict.Search("something else")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		_ = dict.Add("test", "this is just a test")

		want := "this is just a test"
		_, err := dict.Search("test")

		assertError(t, err, nil)
		assertDefinition(t, dict, "test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{}
		definition := "this is just a test"
		_ = dict.Add("test", definition)
		err := dict.Add("test", "overwrite")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, "test", definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"

		dict.Update(word, newDefinition)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "this is just a test"}

	dict.Delete(word)

	_, err := dict.Search(word)

	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dict Dictionary, word, want string) {
	t.Helper()

	got, _ := dict.Search(word)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
