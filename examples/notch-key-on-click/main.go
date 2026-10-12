// Notch-key-on-click shows a screen flyout that takes the keyboard only
// when the user presses in it (FlyoutOptions.KeyOnClick): open a text
// editor behind it, type there, then click the field of the flyout and
// type again. A press outside the flyout closes it.
//
//	go run ./examples/notch-key-on-click
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

type state struct{ text string }

func (s *state) view(c *ui.Context) {
	c.Root().Background(c.Theme().Surface)
	ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
		ui.Text(c, "Typing goes to the app behind until you click the field").Bold()
		ui.TextInput(c.Key("field"), &s.text).Placeholder("Click, then type").Fill()
		ui.Text(c, "A press outside the panel closes it")
	})
}

func main() {
	s := &state{}
	mygo.App.WhenReady(func() {
		d := mygo.Screen.PrimaryDisplay()
		anchor := mygo.Rectangle{X: d.Bounds.X + d.Bounds.Width/2 - 100, Y: d.Bounds.Y, Width: 200, Height: 1}
		fl := mygo.NewFlyout(mygo.FlyoutOptions{
			Screen:     true,
			KeyOnClick: true,
			Anchor:     anchor,
			Placement:  mygo.PlacementBottom,
			Width:      480,
			Height:     160,
			Shadow:     true,
			Content:    ui.View(s.view),
		})
		fl.OnClosed(func() { mygo.App.Quit() })
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
