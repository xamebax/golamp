package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	unicorn "github.com/arussellsaw/unicorn-go"
)

func main() {

	initialBrightness := 45
	c := unicorn.Client{Path: unicorn.SocketPath}

	c.Connect()
	c.Clear()
	c.SetBrightness(uint(initialBrightness))

	// This captures an interrupt and turns all the LEDs off
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		c.Clear()
		c.Show()
		os.Exit(1)
	}()

	for true {
		passing, _ := CheckBuildStatus()

		if passing == false {
			pipelineColor := [3]uint{99, 50, 60}
			pixels := [64]unicorn.Pixel{}
			for i := range pixels {
				pixels[i] = unicorn.Pixel{pipelineColor[0], pipelineColor[1], pipelineColor[2]}
			}
			err := c.SetAllPixels(pixels)
			if err != nil {
				fmt.Println(err)
			}

			Pulsate(c, initialBrightness)

			// slow the loop down to not DDoS Travis
			time.Sleep(2 * time.Second)
		} else {

			c.Clear()
			c.Show()
			// slow the loop down to not DDoS Travis
			time.Sleep(2 * time.Second)
		}
	}
}
