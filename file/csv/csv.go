package csv

import (
	f "encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type csv struct {
	text []string
	opt  Option
}

type Option struct {
	Path string
}

func New(opt Option) *csv {
	return &csv{
		opt: opt,
	}
}

func (c *csv) Add(text string) {
	fmt.Print(text)
	c.text = append(c.text, text)
}

func (c *csv) Write() (path string, err error) {
	path = fmt.Sprint(c.opt.Path, uuid.New().String(), ".csv")
	fmt.Print(path)
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	w := f.NewWriter(file)
	w.Write(c.text)
	w.Flush()
	return
}
