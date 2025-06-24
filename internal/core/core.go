package core

import (
	"log"
	"time"
)

type Core struct {
	timer  *time.Ticker
	events chan []byte
}

func (c *Core) StartLoop() {
	for {
		c.timer = time.NewTicker(1 * time.Second)
		defer c.timer.Stop()

		select {
		case <-c.timer.C:
		case <-c.events:
		}

		log.Println("Core loop is running...")
	}
}
