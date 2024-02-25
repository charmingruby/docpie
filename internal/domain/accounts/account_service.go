package domain

type AuthenticateOutput struct {
	Token string `json:"token"`
}

type AccountUseCase interface {
	Authenticate(email, password string) (*AuthenticateOutput, error)
	Register(name, lastName, email, password string) error
	UploadAvatar() error
}
