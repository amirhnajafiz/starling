package core

import (
	"log"
	"time"
)

type Core struct {
	timer *time.Ticker
}

func (c *Core) StartLoop() {
	for {
		c.timer = time.NewTicker(1 * time.Second)
		defer c.timer.Stop()

		log.Println("Core loop is running...")
	}
}
