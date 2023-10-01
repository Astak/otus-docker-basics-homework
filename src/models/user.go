package models

type User struct {
	ID        int64  `db:"id"`
	UserName  string `db:"username"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}
