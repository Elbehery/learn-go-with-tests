package maps

import "errors"

var ErrKeyNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}
