package app

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	natscfg "github.com/sorawaslocked/ap2final_base/pkg/nats"
	natsconsumer "github.com/sorawaslocked/ap2final_base/pkg/nats/consumer"
	"github.com/sorawaslocked/ap2final_email_service/internal/adapter/mailer"
	"github.com/sorawaslocked/ap2final_email_service/internal/adapter/nats/handler"
	"github.com/sorawaslocked/ap2final_email_service/internal/config"
	"github.com/sorawaslocked/ap2final_email_service/internal/usecase"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const serviceName = "email"

type App struct {
	natsPubSubConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println("starting " + serviceName + " service")

	log.Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))
	natsClient, err := natscfg.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.Nkey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	natsPubSubConsumer := natsconsumer.NewPubSub(natsClient)

	mailerSend := mailer.NewMailer(mailersend.NewMailersend(cfg.MailerKey))

	userUseCase := usecase.NewUser(mailerSend)

	userHandler := handler.NewUser(userUseCase)

	natsPubSubConsumer.Subscribe(natsconsumer.PubSubSubscriptionConfig{
		Subject: cfg.Nats.NatsSubjects.UserEventSubject,
		Handler: userHandler.Handler,
	})

	app := &App{natsPubSubConsumer: natsPubSubConsumer}

	return app, nil
}

func (a *App) Close(_ context.Context) {
	a.natsPubSubConsumer.Stop()
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
	a.natsPubSubConsumer.Start(ctx, errCh)
	log.Println(fmt.Sprintf("service %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Println(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))
		log.Println("graceful shutdown completed!")
	}

	return nil
}
