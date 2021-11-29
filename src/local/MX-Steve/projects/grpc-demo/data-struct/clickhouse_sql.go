package datastruct

import "fmt"

const (
	userSql = `
	SELECT * FROM user WHERE id = ?
	`
)

func getUserInfoQuery(userSql, userId string) string {
	q := fmt.Sprintf(userSql, userId)
	return q
}
