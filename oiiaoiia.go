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
		panic("why")
	}
}

func main() {
	frames := []string{
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
		frame7, frame7,
		frame8, frame8,
		frame9, frame9,
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

func printFrame(frame string) {
	w, h := termbox.Size()
	lines := strings.Split(frame, "\n")
	startY := (h - len(lines)) / 2

	for y, line := range lines {
		startX := (w - len(line)) / 2
		for x, c := range line {
			termbox.SetCell(startX+x, startY+y, c, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
