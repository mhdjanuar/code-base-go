package entities

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) IsValid() bool {
	return u.Email != "" && u.Password != ""
}