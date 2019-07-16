package main

import (
	"fmt"
)

const (
	clearSequence       string = "\033[2J"
	returnToTopSequence string = "\033[H"
)

func terminalExecuteSequence(sequence string) {
	fmt.Println(sequence)
}

func terminalClear() {
	terminalExecuteSequence(clearSequence)
	terminalExecuteSequence(returnToTopSequence)
}
