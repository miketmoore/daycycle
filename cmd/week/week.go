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

type InputHandler interface {
	HandleInput(*pixelgl.Window)
}

type State interface {
	InputHandler
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
	context Day
}

func (s StateDawn) Go() {
	fmt.Printf("StateDawn Go %T\nn", s.context.CurrentState)
	s.context.CurrentState = s.context.States["morning"]
}

func (s StateDawn) HandleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		fmt.Println("StartDawn HandleInput (to StateMorning)")
		s.context.Change("morning")
	}
}

func (s StateDawn) Update(txt *text.Text, win *pixelgl.Window) {
	txt.Clear()
	txt.Dot = txt.Orig
	fmt.Fprintln(txt, s.Name())
	win.Clear(color.RGBA{247, 118, 37, 1})
}

func (s StateDawn) Name() string {
	return "Dawn"
}

// Morning
type StateMorning struct {
	context Day
}

func (s StateMorning) Go() {
	s.context.CurrentState = s.context.States["noon"]
}

func (s StateMorning) HandleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		fmt.Println("StateMorning HandleInput (to StateDawn)")
		s.context.Change("dawn")
	}
}

func (s StateMorning) Name() string {
	return "Morning"
}

func (s StateMorning) Update(txt *text.Text, win *pixelgl.Window) {
	txt.Clear()
	txt.Dot = txt.Orig
	fmt.Fprintln(txt, s.Name())
	win.Clear(color.RGBA{237, 182, 83, 1})
}

// Noon
type StateNoon struct {
	context Day
}

func (s StateNoon) Go() {
	s.context.CurrentState = s.context.States["dawn"]
}

func (s StateNoon) HandleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		fmt.Println("StateNoon HandleInput (to StateDawn)")
		s.context.Change("dawn")
	}
}

func (s StateNoon) Name() string {
	return "Noon"
}

func (s StateNoon) Update(txt *text.Text, win *pixelgl.Window) {
	txt.Clear()
	txt.Dot = txt.Orig
	fmt.Fprintln(txt, s.Name())
	win.Clear(color.RGBA{51, 193, 255, 1})
}

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

	var states = map[string]State{
		"dawn":    StateDawn{},
		"morning": StateMorning{},
		"noon":    StateNoon{},
	}
	day := Day{
		States:       states,
		CurrentState: states["dawn"],
	}
	win.Clear(color.RGBA{7, 15, 35, 1})
	for !win.Closed() {
		txt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(txt.Bounds().Center())))
		fmt.Printf("%T\n", day.CurrentState)
		if fmt.Sprintf("%T", day.CurrentState) == "main.StateDawn" {
			if win.JustPressed(pixelgl.KeyEnter) {
				day.CurrentState.Update(txt, win)
				day.CurrentState = states["morning"]
				// day.CurrentState.Go()
			}
		} else if fmt.Sprintf("%T", day.CurrentState) == "main.StateMorning" {
			if win.JustPressed(pixelgl.KeyEnter) {
				day.CurrentState.Update(txt, win)
				day.CurrentState = states["noon"]
				// day.CurrentState.Go()
			}
		} else if fmt.Sprintf("%T", day.CurrentState) == "main.StateNoon" {
			if win.JustPressed(pixelgl.KeyEnter) {
				day.CurrentState.Update(txt, win)
				day.CurrentState = states["dawn"]
				// day.CurrentState.Go()
			}
		}
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
