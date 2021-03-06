package publisher

import (
	"context"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
)

type Srv interface {
	Connect(ctx context.Context) error
	Send(ev *api.NotificationDTO) error
}
