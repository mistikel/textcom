package csv_test

import (
	"strings"
	"testing"
	"textcom/file/csv"
)

func TestWrite(t *testing.T) {
	opt := csv.Option{
		Path: "/Users/agusmistiawan/Documents/",
	}
	csv := csv.New(opt)
	csv.Add("Test")
	csv.Add("Mock")

	path, err := csv.Write()
	if err != nil {
		t.Errorf("Want: nil got [%v]", err)
	}

	if !strings.Contains(path, "/Users/agusmistiawan/Documents/") {
		t.Errorf("Want: [/Users/agusmistiawan/Documents/] got [%s]", path)
	}
}
