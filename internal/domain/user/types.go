package user

type (
	Role  string
	Login string
	Email string
	Phone string
	Name  string
)

const (
	RoleViewer Role = "viewer"
	RoleAdmin  Role = "admin"
)

func (l *Login) Validate() error {
	if len(*l) < 6 {
		return ErrInvalidLogin
	}
	return nil
}

func (e *Email) Validate() error {
	if len(*e) < 6 {
		return ErrInvalidEmail
	}
	return nil
}

func (p *Phone) Validate() error {
	if len(*p) < 11 {
		return ErrInvalidPhone
	}
	return nil
}

func (n *Name) Validate() error {
	if len(*n) < 2 {
		return ErrInvalidName
	}
	return nil
}
