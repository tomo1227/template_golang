//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/tomo1227/template_golang/internal/adapter/handler"
	"github.com/tomo1227/template_golang/internal/infrastructure/logger"
)

func InjectFiberHandler() handler.FiberHandler {
	wire.Build(
		logger.NewLogger,
		handler.NewFiberHandler,
	)
	return nil
}
