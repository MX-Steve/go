//go:generate mockgen -destination ./mock/user_mock.go -package user_mock -source user.go
package d1

type User interface {
	Name() string
	SetAge(age int) bool
	V(idx int, name string) (string, error)
}
