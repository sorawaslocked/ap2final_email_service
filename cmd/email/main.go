package main

import (
	"context"
	"github.com/sorawaslocked/ap2final_email_service/internal/app"
	"github.com/sorawaslocked/ap2final_email_service/internal/config"
)

func main() {
	ctx := context.Background()

	cfg := config.MustLoad()

	application, err := app.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	err = application.Run()
	if err != nil {
		panic(err)
	}
}
