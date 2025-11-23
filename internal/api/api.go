package api

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type StartAPI interface {
	GetMessage(ctx *th.Context, update telego.Update) error
}
