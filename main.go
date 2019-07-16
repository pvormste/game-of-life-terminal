package main

import (
	"time"
)

const renderThresholdInMilliseconds = 1000 * time.Millisecond

func main() {
	renderTime := time.Date(1, 1, 1970, 0, 0, 0, 0, time.UTC)
	world := NewWorld(30)

	for {
		elapsed := time.Now().Sub(renderTime)
		if elapsed < renderThresholdInMilliseconds {
			continue
		}

		terminalClear()
		world.render()
		world.update()

		renderTime = time.Now()
	}
}
