package search

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you where looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dict Dictionary) Add(word, definition string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExists
	default:
		return nil
	}
	return nil
}

func (dict Dictionary) Update(word, newDefinition string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[word] = newDefinition
	default:
		return nil
	}
	return nil
}

func (dict Dictionary) Delete(word string) {
	delete(dict, word)
}
