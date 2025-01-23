package main

import (
	_ "embed"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

//go:embed ascii/0.txt
var frame0 string

//go:embed ascii/1.txt
var frame1 string

func handleKey(key termbox.Key) {
	switch key {
	case termbox.KeyEsc:
		panic("why")
	}
}

func main() {
	frames := []string{
		frame0,
		frame1,
	}

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	for {
		// Display each frame
		for _, frame := range frames {
			select {
			case e := <-events:
				handleKey(e.Key)
			default:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				printFrame(frame)
				termbox.Flush()
				time.Sleep(100 * time.Millisecond) // Pause for 500 milliseconds
			}
		}
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
