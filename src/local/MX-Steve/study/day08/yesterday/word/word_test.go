package word

import (
	"testing"
)

// test
func TestIsPalindrome(t *testing.T) {
	type test struct {
		str  string
		want bool
	}
	tests := map[string]test{
		"simple":       {"iamai", true},
		"englishFalse": {"abc", false},
		"chineseTrue":  {"上海在海上", true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome(tc.str)
			if got != tc.want {
				t.Errorf("expected:%v,got:%v", tc.want, got)
			}
		})
	}

}
