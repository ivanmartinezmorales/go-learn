package dictionary

import "errors"

// Dictionary class
type Dictionary map[string]string

// ErrNotFound - when a key is not found in the dictionary
var ErrNotFound = errors.New("could not find the word that you were looking for")

// Search - returns a value for a given string keyword
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
