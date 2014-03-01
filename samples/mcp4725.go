package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"

	"github.com/kidoman/embd/controller/mcp4725"
	"github.com/kidoman/embd/i2c"
)

func main() {
	if err := i2c.Open(); err != nil {
		panic(err)
	}
	defer i2c.Close()

	bus := i2c.NewBus(1)

	dac := mcp4725.New(bus, 0x62)
	defer dac.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	for {
		select {
		case <-stop:
			return
		default:
			voltage := rand.Intn(4096)
			if err := dac.SetVoltage(voltage); err != nil {
				log.Printf("mcp4725: %v", err)
			}
		}
	}
}
