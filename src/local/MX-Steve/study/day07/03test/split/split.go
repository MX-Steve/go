package split

import "strings"

// split string
// Split: split string by Sep
func Split(s, sep string) []string {
	count := strings.Count(s, sep) // contains the count of sep in s
	result := make([]string, 0, count+1)
	index := strings.Index(s, sep)
	for index >= 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}

// Add
func Add(a, b int) (sum int) {
	return a + b
}
