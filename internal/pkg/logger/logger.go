package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

// New creates a new logger
func New() (*Logger, error) {
	var err error
	l := &Logger{}

	l.Logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return l, nil
}
