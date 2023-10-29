package maps

import "errors"

var ErrKeyNotFound = errors.New("could not find the word you were looking for")
var ErrKeyAlreadyExist = errors.New("cannot add word because it already exists")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
		return nil
	}
	return ErrKeyAlreadyExist
}
