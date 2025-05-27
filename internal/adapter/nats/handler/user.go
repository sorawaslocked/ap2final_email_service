package handler

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/sorawaslocked/ap2final_email_service/internal/adapter/nats/handler/dto"
	"log"
)

type User struct {
	useCase UserUseCase
}

func NewUser(useCase UserUseCase) *User {
	return &User{
		useCase: useCase,
	}
}

func (u *User) Handler(ctx context.Context, msg *nats.Msg) error {
	user, err := dto.FromRegisterEventToUser(msg)
	if err != nil {
		log.Println("failed to convert Client NATS msg", err)

		return err
	}

	err = u.useCase.Send(ctx, user)
	if err != nil {
		log.Println("failed to send user", err)

		return err
	}

	return nil
}
