package file

import (
	"errors"
	"textcom/file/csv"
)

const (
	CSV = iota + 1
)

type IFile interface {
	Add(text string)
	Write() (path string, err error)
}

func New(_type int) (IFile, error) {
	var f IFile
	switch _type {
	case CSV:
		f = csv.New()
	default:
		return nil, errors.New("undefined type")
	}

	return f, nil
}
