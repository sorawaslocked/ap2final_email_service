package config

import "github.com/sorawaslocked/ap2final_base/pkg/nats"

type (
	Config struct {
		MailerKey string      `yaml:"mailerKey" env-required:"true"`
		Nats      nats.Config `yaml:"nats" env-required:"true"`
	}
)
