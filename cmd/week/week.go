package main

import (
	"bufio"
	"flag"
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

var tmpFile string = "/tmp/week"
var debug bool = false

var locale = map[string]string{
	"gameTitle": "One Week",
}

func getFlags() bool {
	d := flag.Bool("d", false, "enable debug mode, which logs to a temporary file")
	flag.Parse()
	return *d
}

// pixel "main"
func run() {
	debug = getFlags()

	// Setup a logger
	writer, logger := initLogger()
	logger.Print("run")
	writer.Flush()

	// Setup fonts
	txt := initText()
	fmt.Fprintln(txt, locale["gameTitle"])

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
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
		win.Update()
		writer.Flush()
	}
}

// pixel takes over the main function
func main() {
	pixelgl.Run(run)
}

func initLogger() (*bufio.Writer, *log.Logger) {
	f, err := os.Create(tmpFile)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	logger := log.New(writer, "INFO: ", log.Lshortfile)
	return writer, logger
}

func initText() *text.Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	orig := pixel.V(20, 50)
	txt := text.New(orig, atlas)
	txt.Color = colornames.Darkslategrey
	return txt
}
