package handler

import (
	"context"
	"github.com/sorawaslocked/ap2final_email_service/internal/model"
)

type UserUseCase interface {
	Send(ctx context.Context, user model.User) error
}
