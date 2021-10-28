// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

//go:generate mockgen -destination mock_concurrent_test.go -package concurrent_test  local/MX-Steve/study/others/gomock/demo02/concurrent Math

type Math interface {
	Sum(a, b int) int
}

func GetSum(m Math) int {
	return m.Sum(5, 6)
}
