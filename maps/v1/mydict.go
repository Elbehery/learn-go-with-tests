package main

import (
	"errors"
)

var ErrNoMatch = errors.New("key not found in MyDict")

type MyDict map[string]string

func NewMyDict() MyDict {
	return MyDict{}
}

func (d MyDict) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNoMatch
	}
	return value, nil
}
