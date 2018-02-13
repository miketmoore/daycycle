package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
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

var locale = map[string]string{
	"gameTitle": "Day Cycle",
}

var palette = map[string]color.RGBA{
	"dawn":      color.RGBA{247, 118, 37, 1},
	"morning":   color.RGBA{237, 182, 83, 1},
	"noon":      color.RGBA{51, 193, 255, 1},
	"afternoon": color.RGBA{51, 193, 255, 1},
	"dusk":      color.RGBA{92, 56, 214, 1},
	"evening":   color.RGBA{51, 193, 255, 1},
	"night":     color.RGBA{74, 58, 94, 1},
	"midnight":  color.RGBA{30, 24, 38, 1},
}

var tmpFile string = "/tmp/week"
var debug bool = false

type State interface {
	Name() string
	Update(*text.Text, *pixelgl.Window)
	Go()
}

// FSM Parts of a Day
type Day struct {
	States       map[string]State
	CurrentState State
}

func (d *Day) Change(stateKey string) {
	fmt.Printf("Day Change to %s\n", stateKey)
	fmt.Printf("0 %T\n", d.CurrentState)
	d.CurrentState = d.States[stateKey]
	fmt.Printf("1 %T\n", d.CurrentState)
}

func (d *Day) Start(stateKey string) {
	fmt.Printf("Day Start stateKey: %s\n", stateKey)
	d.CurrentState = d.States[stateKey]
	fmt.Printf("Day Start CurrentState type: %T\n", d.CurrentState)
}

// Dawn
type StateDawn struct {
	context *Day
}

func (s StateDawn) Go() {
	fmt.Printf("StateDawn Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["morning"]
}

func (s StateDawn) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["dawn"])
	}
}

func (s StateDawn) Name() string {
	return "Dawn"
}

// Morning
type StateMorning struct {
	context *Day
}

func (s StateMorning) Go() {
	fmt.Printf("StateMorning Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["noon"]
}

func (s StateMorning) Name() string {
	return "Morning"
}

func (s StateMorning) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["morning"])
	}
}

// Noon
type StateNoon struct {
	context *Day
}

func (s StateNoon) Go() {
	fmt.Printf("StateNoon Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["afternoon"]
}

func (s StateNoon) Name() string {
	return "Noon"
}

func (s StateNoon) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["noon"])
	}
}

// Afternoon
type StateAfternoon struct {
	context *Day
}

func (s StateAfternoon) Go() {
	fmt.Printf("StateAfternoon Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["dusk"]
}

func (s StateAfternoon) Name() string {
	return "Afternoon"
}

func (s StateAfternoon) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["afternoon"])
	}
}

// Dusk
type StateDusk struct {
	context *Day
}

func (s StateDusk) Go() {
	fmt.Printf("StateDusk Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["evening"]
}

func (s StateDusk) Name() string {
	return "Dusk"
}

func (s StateDusk) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["dusk"])
	}
}

// Evening
type StateEvening struct {
	context *Day
}

func (s StateEvening) Go() {
	fmt.Printf("StateEvening Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["night"]
}

func (s StateEvening) Name() string {
	return "Evening"
}

func (s StateEvening) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["evening"])
	}
}

// Night
type StateNight struct {
	context *Day
}

func (s StateNight) Go() {
	fmt.Printf("StateNight Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["midnight"]
}

func (s StateNight) Name() string {
	return "Night"
}

func (s StateNight) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		txt.Color = colornames.White
		win.Clear(palette["night"])
	}
}

// Midnight
type StateMidnight struct {
	context *Day
}

func (s StateMidnight) Go() {
	fmt.Printf("StateMidnight Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["dawn"]
}

func (s StateMidnight) Name() string {
	return "Midnight"
}

func (s StateMidnight) Update(txt *text.Text, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		s.Go()
		txt.Clear()
		txt.Dot = txt.Orig
		fmt.Fprintln(txt, s.Name())
		win.Clear(palette["midnight"])
	}
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

	// Initialize context "Day"
	var day = &Day{}

	// Initialize states and assign to context
	day.States = map[string]State{
		"dawn":      StateDawn{day},
		"morning":   StateMorning{day},
		"noon":      StateNoon{day},
		"afternoon": StateAfternoon{day},
		"dusk":      StateDusk{day},
		"evening":   StateEvening{day},
		"night":     StateNight{day},
		"midnight":  StateMidnight{day},
	}
	day.CurrentState = day.States["dawn"]

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
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	orig := pixel.V(20, 50)
	txt := text.New(orig, atlas)
	txt.Color = colornames.White
	return txt
}
