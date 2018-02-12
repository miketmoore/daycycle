package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

// Week is an implementation of a finite state machine in Go.
// Each state represents a section of time during the week.
// I guess this is a type of clock.

// https://english.stackexchange.com/questions/28498/precise-names-for-parts-of-a-day

type InputHandler interface {
	HandleInput(pixelgl.Window)
}

// FSM Weekdays
// Sunday
// Monday
// Tuesday
// Wendesday
// Thursday
// Friday
// Saturday

// type Weekday int

// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wendesday
// 	Thursday
// 	Friday
// 	Saturday
// )

type StateSunday struct {
}

func (s StateSunday) HandleInput(win *pixelgl.Window) {
}

// FSM Parts of a Day
// Dawn
// Morning
// Noon
// Afternoon
// Dusk
// Evening
// Night
// Midnight

var locale = map[string]string{
	"gameTitle": "One Week",
}

func run() {
	// Setup a logger
	writer, logger := initLogger()
	logger.Print("run")
	writer.Flush()

	// Setup fonts
	displayTxt := initDisplayText()
	fmt.Fprintln(displayTxt, locale["gameTitle"])

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  locale["gameTitle"],
		Bounds: pixel.R(0, 0, 400, 225),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Antiquewhite)
		displayTxt.Draw(win, pixel.IM.Scaled(displayTxt.Orig, 2))
		win.Update()
		writer.Flush()
	}
}

func main() {
	pixelgl.Run(run)
}

func initLogger() (*bufio.Writer, *log.Logger) {
	f, err := os.Create("/tmp/dat2")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	logger := log.New(writer, "INFO: ", log.Lshortfile)
	return writer, logger
}

func initDisplayText() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	orig := pixel.V(20, 20)
	txt := text.New(orig, atlas)
	txt.Color = colornames.Darkslategrey
	return txt
}
