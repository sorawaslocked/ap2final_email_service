package dto

import (
	"github.com/nats-io/nats.go"
	"github.com/sorawaslocked/ap2final_email_service/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/events"
	"google.golang.org/protobuf/proto"
)

func FromRegisterEventToUser(msg *nats.Msg) (model.User, error) {
	var pbUser events.UserRegisterEvent
	err := proto.Unmarshal(msg.Data, &pbUser)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:    pbUser.UserID,
		Email: pbUser.Email,
	}, nil
}
