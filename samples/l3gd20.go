package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/kid0m4n/go-rpi/i2c"
	"github.com/kid0m4n/go-rpi/sensor/l3gd20"
)

func main() {
	bus, err := i2c.NewBus(1)
	if err != nil {
		log.Panic(err)
	}
	gyro := l3gd20.New(bus, l3gd20.R250DPS)
	defer gyro.Close()

	gyro.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	orientations, err := gyro.Orientations()
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case orientation := <-orientations:
			log.Printf("x: %v, y: %v, z: %v", orientation.X, orientation.Y, orientation.Z)
		case <-quit:
			return
		}
	}
}
