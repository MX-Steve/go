package word

import "unicode"

// 回文测试
// IsPalindrome 回文测试
func IsPalindrome(s string) bool {
	var letters []rune
	for _, l := range s {
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))
		}
	}
	for i := 0; i < len(letters); i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
