package domain

import "errors"

type Password string

func (p *Password) Validate() error {
	if len(*p) < 6 {
		return errors.New("password_too_weak")
	}
	return nil
}
