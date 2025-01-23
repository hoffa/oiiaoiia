package main

import (
	_ "embed"
	"log"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

//go:embed ascii/0.txt
var frame0 string

//go:embed ascii/1.txt
var frame1 string

func main() {
	frames := []string{
		frame0,
		frame1,
	}

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	// Display each frame
	for _, frame := range frames {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		//println(frame)
		printFrame(frame)
		termbox.Flush()
		time.Sleep(1000 * time.Millisecond) // Pause for 500 milliseconds
	}
}

func printFrame(frame string) {
	// Get the size of the terminal
	width, height := termbox.Size()

	// Split the frame into lines
	lines := splitLines(frame)

	// Calculate the starting position to center the frame
	startY := (height - len(lines)) / 2

	for y, line := range lines {
		startX := (width - len(line)) / 2
		for x, c := range line {
			termbox.SetCell(startX+x, startY+y, c, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}
