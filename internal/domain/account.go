package domain

import "time"

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	AvatarURL string    `json:"avatar_url"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthenticateOutput struct {
	Token string `json:"token"`
}

type AccountUseCase interface {
	Authenticate(email, password string) (*AuthenticateOutput, error)
	Register(name, lastName, email, password string) error
	UploadAvatar() error
}
