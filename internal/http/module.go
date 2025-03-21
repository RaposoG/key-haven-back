package http

import (
	"key-haven-back/internal/handler"
	"key-haven-back/internal/router"
	"key-haven-back/pkg/docs"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
	handler.Module,
	router.Module,
	docs.Module,
)
