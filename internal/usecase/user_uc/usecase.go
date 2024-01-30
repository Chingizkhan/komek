package user_uc

type UseCase struct {
	r UserRepository
}

func New(r UserRepository) *UseCase {
	return &UseCase{r}
}

func (u *UseCase) Register() error {
	return nil
}

func (u *UseCase) Login() error {
	return nil
}

func (u *UseCase) Logout() error {
	return nil
}

func (u *UseCase) Delete() error {
	return nil
}

func (u *UseCase) ChangePassword() error {
	return nil
}

func (u *UseCase) Update() error {
	return nil
}
