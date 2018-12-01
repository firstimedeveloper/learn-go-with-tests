package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("The word already exists")
var ErrWordNotExists = errors.New("The word doesn't exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, def string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}
	d[word] = def

	return nil
}

func (d Dictionary) Update(word, newDef string) error {
	if _, ok := d[word]; ok {
		d[word] = newDef
		return nil
	}
	return ErrWordNotExists
}
