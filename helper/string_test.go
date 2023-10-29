package helper

import (
	"testing"
)

func TestIterator(t *testing.T) {
	stringCollection := NewStringCollection("testing")
	it := stringCollection.createIterator()
	i := 0
	for it.hasNext() {
		text := it.getNext()
		switch i {
		case 0:
			if text != "t" {
				t.Errorf("Want: t got [%s]", text)
			}
		case 1:
			if text != "e" {
				t.Errorf("Want: e got [%s]", text)
			}
		case 2:
			if text != "s" {
				t.Errorf("Want: s got [%s]", text)
			}
		case 3:
			if text != "t" {
				t.Errorf("Want: t got [%s]", text)
			}
		case 4:
			if text != "i" {
				t.Errorf("Want: i got [%s]", text)
			}
		case 5:
			if text != "n" {
				t.Errorf("Want: n got [%s]", text)
			}
		case 6:
			if text != "g" {
				t.Errorf("Want: g got [%s]", text)
			}
		}
		i++
	}

	if it.hasNext() {
		t.Errorf("Want: false got true")
	}
}
