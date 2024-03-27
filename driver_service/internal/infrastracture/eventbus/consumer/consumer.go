package consumer

import "context"

type Consumer interface {
	Consume(ctx context.Context)
}
