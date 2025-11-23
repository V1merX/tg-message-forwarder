package main

import (
	"context"
	"log"

	"github.com/V1merX/tg-message-forwarder/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.New()

	if err := a.Run(ctx); err != nil {
		log.Fatal("Failed to start app")
	}
}
