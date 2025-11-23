package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/V1merX/tg-message-forwarder/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.New()

	if err := a.Run(ctx); err != nil {
		slog.Error("Failed to start app", slog.Any("error", err))
		os.Exit(1)
	}
}
