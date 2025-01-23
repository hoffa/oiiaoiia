package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	// Directory containing ASCII art frames
	dir := "./ascii"

	// Read all files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Array to store frames
	var frames []string

	// Read each file and store its content in the frames array
	for _, file := range files {
		if !file.IsDir() {
			if filepath.Ext(file.Name()) != ".txt" {
				continue
			}
			content, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			frames = append(frames, string(content))
		}
	}

	// Display each frame
	for _, frame := range frames {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		//println(frame)
		printFrame(frame)
		termbox.Flush()
		time.Sleep(10 * time.Millisecond) // Pause for 500 milliseconds
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
