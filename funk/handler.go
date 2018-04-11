package funk

import (
	"time"

	"github.com/google/uuid"
)

type EventProcessor interface {
	Process() error
	Handler() func(Event)
}

func NewGeneratorEventProcessor(handler func(Event)) EventProcessor {
	return generatorEventProcessor{
		handler: handler,
	}
}

type generatorEventProcessor struct {
	handler func(Event)
}

func (gep generatorEventProcessor) Process() error {
	for {
		evt := Event{
			Id:        uuid.New().String(),
			Timestamp: time.Now().UnixNano(),
			Src:       "local",
			Data:      []byte{1, 2, 3, 4, 5},
		}
		gep.Handler()(evt)
		time.Sleep(1 * time.Second)
	}
}

func (gep generatorEventProcessor) Handler() func(Event) {
	return gep.handler
}
