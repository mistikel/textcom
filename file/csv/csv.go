package csv

import (
	f "encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const (
	prefixPath = "/Users/agusmistiawan/Documents/%s.csv"
)

type csv struct {
	text []string
}

func New() *csv {
	return &csv{}
}

func (c *csv) Add(text string) {
	fmt.Print(text)
	c.text = append(c.text, text)
}

func (c *csv) Write() (path string, err error) {
	path = fmt.Sprintf(prefixPath, uuid.New().String())
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
