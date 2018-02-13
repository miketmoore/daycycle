package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

// "Day Cycle" is an implementation of the state pattern in Go.
// Each state represents a section of time during the week.
// I guess this is a type of clock.

// https://english.stackexchange.com/questions/28498/precise-names-for-parts-of-a-day

// i18n
var locale = map[string]string{
	"title":     "day cycle",
	"subtitle":  "press enter",
	"dawn":      "dawn",
	"morning":   "morning",
	"noon":      "noon",
	"afternoon": "afternoon",
	"dusk":      "dusk",
	"evening":   "evening",
	"night":     "night",
	"midnight":  "midnight",
}

type StateName string

const (
	Dawn      StateName = "dawn"
	Morning   StateName = "morning"
	Noon      StateName = "noon"
	Afternoon StateName = "afternoon"
	Dusk      StateName = "dusk"
	Evening   StateName = "evening"
	Night     StateName = "night"
	Midnight  StateName = "midnight"
)

var palette = map[StateName]color.RGBA{
	Dawn:      color.RGBA{25, 26, 21, 1},
	Morning:   color.RGBA{73, 61, 63, 1},
	Noon:      color.RGBA{105, 94, 88, 1},
	Afternoon: color.RGBA{143, 133, 124, 1},
	Dusk:      color.RGBA{105, 94, 88, 1},
	Evening:   color.RGBA{73, 61, 63, 1},
	Night:     color.RGBA{25, 26, 21, 1},
	Midnight:  colornames.Black,
}

var tmpFile string = "/tmp/week"
var debug bool = false
var writer *bufio.Writer
var logger *log.Logger

type State interface {
	Name() string
	Update(*text.Text, *pixelgl.Window)
	Go()
}

// FSM Parts of a Day
type Day struct {
	States       map[StateName]State
	CurrentState State
}

// Dawn
type StateDawn struct {
	context *Day
}

func (s StateDawn) Go() {
	s.context.CurrentState = s.context.States[Morning]
}

func (s StateDawn) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Dawn])
	}
}

func (s StateDawn) Name() string {
	return strings.Title(locale["dawn"])
}

// Morning
type StateMorning struct {
	context *Day
}

func (s StateMorning) Go() {
	s.context.CurrentState = s.context.States[Noon]
}

func (s StateMorning) Name() string {
	return strings.Title(locale["morning"])
}

func (s StateMorning) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Morning])
	}
}

// Noon
type StateNoon struct {
	context *Day
}

func (s StateNoon) Go() {
	s.context.CurrentState = s.context.States[Afternoon]
}

func (s StateNoon) Name() string {
	return strings.Title(locale["noon"])
}

func (s StateNoon) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Noon])
	}
}

// Afternoon
type StateAfternoon struct {
	context *Day
}

func (s StateAfternoon) Go() {
	s.context.CurrentState = s.context.States[Dusk]
}

func (s StateAfternoon) Name() string {
	return strings.Title(locale["afternoon"])
}

func (s StateAfternoon) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Afternoon])
	}
}

// Dusk
type StateDusk struct {
	context *Day
}

func (s StateDusk) Go() {
	s.context.CurrentState = s.context.States[Evening]
}

func (s StateDusk) Name() string {
	return strings.Title(locale["dusk"])
}

func (s StateDusk) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Dusk])
	}
}

// Evening
type StateEvening struct {
	context *Day
}

func (s StateEvening) Go() {
	s.context.CurrentState = s.context.States[Night]
}

func (s StateEvening) Name() string {
	return strings.Title(locale["evening"])
}

func (s StateEvening) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Evening])
	}
}

// Night
type StateNight struct {
	context *Day
}

func (s StateNight) Go() {
	s.context.CurrentState = s.context.States[Midnight]
}

func (s StateNight) Name() string {
	return strings.Title(locale["night"])
}

func (s StateNight) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette[Night])
	}
}

// Midnight
type StateMidnight struct {
	context *Day
}

func (s StateMidnight) Go() {
	s.context.CurrentState = s.context.States[Dawn]
}

func (s StateMidnight) Name() string {
	return strings.Title(locale["midnight"])
}

func (s StateMidnight) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		fmt.Fprintln(txt, s.Name())
		win.Clear(palette[Midnight])
	}
}

func getFlags() bool {
	d := flag.Bool("d", false, "enable debug mode, which logs to a temporary file")
	flag.Parse()
	return *d
}

func d(s ...interface{}) {
	if debug == true {
		logger.Print(s)
	}
}

// pixel "main"
func run() {
	debug = getFlags()

	// Setup a logger
	writer, logger = initLogger()
	d("run")
	writer.Flush()

	// Setup fonts
	txt := initText()
	lines := []string{strings.Title(locale["title"]), strings.Title(locale["subtitle"])}
	for _, line := range lines {
		txt.Dot.X -= txt.BoundsOf(line).W() / 2
		fmt.Fprintln(txt, line)
	}

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  locale["title"],
		Bounds: pixel.R(0, 0, 400, 225),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize context "Day"
	var day = &Day{}

	// Initialize states and assign to context
	day.States = map[StateName]State{
		Dawn:      StateDawn{day},
		Morning:   StateMorning{day},
		Noon:      StateNoon{day},
		Afternoon: StateAfternoon{day},
		Dusk:      StateDusk{day},
		Evening:   StateEvening{day},
		Night:     StateNight{day},
		Midnight:  StateMidnight{day},
	}
	day.CurrentState = day.States[Dawn]

	for !win.Closed() {
		txt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(txt.Bounds().Center())))
		day.CurrentState.Update(txt, win)
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
	orig := pixel.V(20, 50)
	txt := text.New(orig, text.Atlas7x13)
	txt.Color = colornames.White
	return txt
}
