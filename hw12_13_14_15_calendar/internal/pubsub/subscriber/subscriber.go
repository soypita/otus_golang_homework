package subscriber

import (
	"context"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
)

type Srv interface {
	Listen(context.Context, func(*api.NotificationDTO) error) error
}
