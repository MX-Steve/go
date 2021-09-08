package split

import "strings"

// split string
// Split: split string by Sep
func Split(s, sep string) (result []string) {
	index := strings.Index(s, sep)
	for index >= 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

// Add
func Add(a, b int) (sum int) {
	return a + b
}
