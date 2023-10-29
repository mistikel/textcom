package helper

import "strings"

type Collection interface {
	createIterator() Iterator
}

type Iterator interface {
	hasNext() bool
	getNext() string
}

type stringIterator struct {
	index int
	text  []string
}

type StringCollection struct {
	text string
}

func (s *StringCollection) createIterator() Iterator {
	return &stringIterator{
		text: strings.Split(s.text, ""),
	}
}

func (s *stringIterator) hasNext() bool {
	return s.index < len(s.text)
}

func (s *stringIterator) getNext() string {
	if s.hasNext() {
		t := s.text[s.index]
		s.index++
		return t
	}
	return ""
}

func NewStringCollection(s string) StringCollection {
	return StringCollection{
		text: s,
	}
}
