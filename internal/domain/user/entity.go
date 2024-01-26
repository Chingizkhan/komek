package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id uuid.UUID
	Name
	Login
	Email
	Phone
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id uuid.UUID, login, email, phone, name string) User {
	return User{
		Id:    id,
		Name:  Name(name),
		Login: Login(login),
		Email: Email(email),
		Phone: Phone(phone),
	}
}

func (u *User) Validate() error {
	if err := u.Name.Validate(); err != nil {
		return err
	}
	if err := u.Login.Validate(); err != nil {
		return err
	}
	if err := u.Email.Validate(); err != nil {
		return err
	}
	if err := u.Phone.Validate(); err != nil {
		return err
	}
	return nil
}
