package interfaces

import "consumer-sample/internal/infrastructure/broker/models"

type Consumer interface {
	Open() error
	Close() error
	Subscribe(models.Subscription)
	Unsubscribe(models.Subscription)
}
