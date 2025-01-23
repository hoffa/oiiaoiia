package main

import (
	_ "embed"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

const frameWidth = 80
const frameHeight = 40

//go:embed ascii/0.txt
var frame0 string

//go:embed ascii/1.txt
var frame1 string

//go:embed ascii/2.txt
var frame2 string

//go:embed ascii/3.txt
var frame3 string

//go:embed ascii/4.txt
var frame4 string

//go:embed ascii/5.txt
var frame5 string

//go:embed ascii/6.txt
var frame6 string

//go:embed ascii/7.txt
var frame7 string

//go:embed ascii/8.txt
var frame8 string

//go:embed ascii/9.txt
var frame9 string

func handleKey(key termbox.Key) {
	switch key {
	case termbox.KeyEsc:
		panic(":(")
	}
}

func printFrame(frame string) {
	w, h := termbox.Size()
	startX := (w - frameWidth) / 2
	startY := (h - frameHeight) / 2

	lines := strings.Split(frame, "\n")
	for y, line := range lines {
		for x, c := range line {
			termbox.SetCell(startX+x, startY+y, c, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func main() {
	frames := []string{
		// Stare 1
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,

		// Spin fast 1
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,

		// Stare 2
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,

		// Spin slow 1
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,

		// Stare 3
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,
		frame0,

		// Spin slow 2
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
		frame1, frame1,
		frame2, frame2,
		frame3, frame3,
		frame4, frame4,
		frame5, frame5,
		frame6, frame6,

		// Freeze butt
		frame6,
		frame6,
		frame6,
		frame6,
		frame6,
		frame6,
		frame6,

		// Spin fast 2
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
	}

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
		for _, frame := range frames {
			select {
			case e := <-events:
				handleKey(e.Key)
			default:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				printFrame(frame)
				termbox.Flush()
				time.Sleep(40 * time.Millisecond)
			}
		}
	}
}
