package repository

type IUser interface {
	Register(username, password string) error
	// RegisterWithTimestamp(username string, password string, createdAt time,Time) error
}

type User struct {
}

func NewUser() IUser {
	return &User{}
}

func (u *User) Register(username, password string) error {
	return nil
}
