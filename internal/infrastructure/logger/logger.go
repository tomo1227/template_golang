package logger

import (
	"log/slog"
	"os"

	"github.com/samber/do"
)

func NewLogger(i *do.Injector) (*slog.Logger, error) {
	return slog.New(slog.NewTextHandler(os.Stdout, nil)), nil
}
