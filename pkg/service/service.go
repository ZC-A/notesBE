package service

import (
	"context"
)

type Service interface {
	Type() string
	Start(context.Context)
	Close()
	Reload(context.Context)
	Wait()
}
