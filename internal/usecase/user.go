package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_email_service/internal/model"
)

type User struct {
	sender EmailPresenter
}

func NewUser(sender EmailPresenter) *User {
	return &User{
		sender: sender,
	}
}

func (u *User) Send(ctx context.Context, user model.User) error {
	return u.sender.Send(ctx, user)
}
