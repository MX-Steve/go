package user

type User struct {
	Name string
}

type UserRepository interface {
	FindOne(id int) (User, error)
}
