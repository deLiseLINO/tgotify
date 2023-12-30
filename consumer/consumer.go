package consumer

import (
	"log"
	telegram "tgotify/client"
	"time"
)

const pollingInterval = 50 // ms

type Fetcher interface {
	Fetch(limit int) ([]telegram.Update, error)
}

type Processor interface {
	ProcessMessage(upd telegram.Update) error
}

type Consumer struct {
	fetcher   Fetcher
	processor Processor
	batchSize int
}

func New(fetcher Fetcher, processor Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(pollingInterval * time.Millisecond)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c *Consumer) handleEvents(events []telegram.Update) error {
	for _, event := range events {
		if err := c.processor.ProcessMessage(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())
			continue
		}
	}

	return nil
}
