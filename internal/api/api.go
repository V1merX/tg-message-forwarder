package api

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type StartAPI interface {
	GetMessage(ctx *th.Context, update telego.Update) error
	MessagePulling(ctx context.Context, bot *telego.Bot, adminID int64) error
}
