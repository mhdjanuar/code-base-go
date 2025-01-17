package entities

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) IsValid() bool {
	return u.Email != "" && u.Password != ""
}