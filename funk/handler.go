package funk

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EventProcessor interface {
	Process() error
	Handler() func(Event) ReturnValue
}

func NewGeneratorEventProcessor(handler func(Event) ReturnValue) EventProcessor {
	return generatorEventProcessor{
		handler: handler,
	}
}

type generatorEventProcessor struct {
	handler func(Event) ReturnValue
}

func (gep generatorEventProcessor) Process() error {
	for {
		evt := Event{
			Id:        uuid.New().String(),
			Timestamp: time.Now().UnixNano(),
			Src:       "local",
			Data:      []byte{1, 2, 3, 4, 5},
		}
		result := gep.Handler()(evt)
		fmt.Printf("Result: %v\n", result)
		time.Sleep(1 * time.Second)
	}
}

func (gep generatorEventProcessor) Handler() func(Event) ReturnValue {
	return gep.handler
}
