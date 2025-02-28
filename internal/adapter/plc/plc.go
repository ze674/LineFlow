package plc

import "context"

type Plc interface {
	HandleProductSignal(ctx context.Context) (<-chan struct{}, error)
	RejectorOn() error
	RejectorOff() error
}
