package helper

import "strings"

// UpperCase simple string convertion
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// LowerCase lower case string
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// CamelCase
// using iterator to iterate list of string and convert to Upper / Lower
func CamelCase(s string) string {
	stringCollection := NewStringCollection(s)
	iterator := stringCollection.createIterator()
	var sb strings.Builder
	isUpperCase := false
	for iterator.hasNext() {
		text := iterator.getNext()
		text = UpperCase(text)
		if !isUpperCase {
			text = LowerCase(text)
		}
		isUpperCase = !isUpperCase
		sb.WriteString(text)
	}
	return sb.String()
}
