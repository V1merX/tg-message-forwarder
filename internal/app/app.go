package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	diContainer *diContainer
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.initDeps(ctx); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("Shutting down app...")

	// TODO: create gracefull shut down deps

	slog.Info("App stopped")
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func() error{
		a.initDI,
		func() error {
			err := a.startBot(ctx)
			return err
		},
	}

	for _, step := range initFuncs {
		if err := step(); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDI() error {
	a.diContainer = NewDIContainer()
	return nil
}

func (a *App) startBot(ctx context.Context) error {
	_, err := a.diContainer.InitKafka(ctx)
	if err != nil {
		panic(err)
	}

	bot, err := a.diContainer.InitBot()
	if err != nil {
		panic(err)
	}

	if err := bot.Start(ctx); err != nil {
		panic(err)
	}

	slog.Info("connected to kafka")
	return nil
}
