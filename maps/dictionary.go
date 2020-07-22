package dictionary

// Dictionary (map) type
type Dictionary map[string]string

const (
	// ErrNotFound - error message
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists - error message
	ErrWordExists = DictionaryErr("that word already exists with a definition")
	// ErrWordDoesNotExist - error message
	ErrWordDoesNotExist = DictionaryErr("that word doesn't exist")
)

// DictionaryErr - more explicit type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search method on Dictionary
func (d Dictionary) Search(searchTerm string) (string, error) {
	definition, ok := d[searchTerm]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add method to allow users to add words
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update - update an existing word in the Dictionary or return an error if it doesn't exist
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = definition
		return nil
	}
	return ErrWordDoesNotExist
}

// Delete - send and forget, no need to return errors
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Search - takes a map and a searchTerm and returns the value associated with the key(searchTerm)
func Search(dictionary Dictionary, searchTerm string) string {
	return dictionary[searchTerm]
}
