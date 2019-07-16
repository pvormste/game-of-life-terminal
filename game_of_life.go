package main

import (
	"bytes"
	"fmt"
)

const (
	alive = true
	dead  = false
)

const (
	livingInhabitant = "o"
	deadInhabitant   = ""
)

type World struct {
	inhabitants [][]bool
	size        int
}

func NewWorld(size int) World {
	inhabitants := make([][]bool, size)
	for y := range inhabitants {
		inhabitants[y] = make([]bool, size)
		for x := 0; x < len(inhabitants[y]); x++ {
			inhabitants[y][x] = true
		}
	}

	return World{
		inhabitants: inhabitants,
		size:        size,
	}
}

func (w World) prepareNextGeneration() [][]bool {
	nextGeneration := make([][]bool, w.size)
	for y := range w.inhabitants {
		nextGeneration[y] = make([]bool, w.size)
		for x := range w.inhabitants[y] {
			nextGeneration[y][x] = w.inhabitants[y][x]
		}
	}

	return nextGeneration
}

func (w *World) update() {
	nextGeneration := w.prepareNextGeneration()
	for y := range w.inhabitants {
		for x := range w.inhabitants[y] {
			w.updateInhabitant(x, y, nextGeneration)
		}
	}

	w.inhabitants = nextGeneration
}

func (w World) updateInhabitant(x, y int, nextGeneration [][]bool) {
	livingNeighbors := w.countLivingNeighbors(x, y)

	currentInhabitantIsAlive := w.inhabitants[y][x]
	if currentInhabitantIsAlive {
		w.updateLivingInhabitant(x, y, livingNeighbors, nextGeneration)
	} else {
		w.updateDeadInhabitant(x, y, livingNeighbors, nextGeneration)
	}
}

func (w World) countLivingNeighbors(x, y int) int {
	livingNeighbors := 0

	xOffsets := []int{-1, 0, 1}
	yOffsets := []int{-1, 0, 1}

	for yOffset := range yOffsets {
		for xOffset := range xOffsets {
			if xOffset == 0 && yOffset == 0 {
				continue
			}

			neighborIsAlive := w.checkNeighborIsAlive(xOffset, yOffset)
			if neighborIsAlive {
				livingNeighbors++
			}
		}
	}

	return livingNeighbors
}

func (World) updateDeadInhabitant(x, y, livingNeighbors int, nextGeneration [][]bool) {
	if livingNeighbors == 3 {
		nextGeneration[y][x] = true
	}
}

func (World) updateLivingInhabitant(x, y, livingNeighbors int, nextGeneration [][]bool) {
	if livingNeighbors < 2 || livingNeighbors > 3 {
		nextGeneration[y][x] = false
	}
}

func (w World) checkNeighborIsAlive(x, y int) bool {
	if x < 0 || y < 0 || x >= w.size || y >= w.size {
		return false
	}

	return w.inhabitants[y][x]
}

func (w World) render() {
	inhabitantsBuffer := w.writeToBuffer()
	fmt.Println(inhabitantsBuffer.String())
}

func (w World) writeToBuffer() *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	for y := range w.inhabitants {
		for x := range w.inhabitants[y] {
			inhabitantCharacter := w.writeInhabitantAsCharacterBytes(w.inhabitants[y][x])
			buffer.Write(inhabitantCharacter)
		}
		buffer.WriteString("\n")
	}

	return buffer
}

func (World) writeInhabitantAsCharacterBytes(isAlive bool) []byte {
	if !isAlive {
		return []byte(deadInhabitant)
	}

	return []byte(livingInhabitant)
}
