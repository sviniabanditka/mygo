// Notch-screen-notch prints the notch of the built-in display, when it has
// one, and hangs a flyout from it, or from the top center of the primary
// display on a Mac without a notch.
//
//	go run ./examples/notch-screen-notch
package main

import (
	"fmt"
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

func main() {
	mygo.App.WhenReady(func() {
		anchor, ok := mygo.Screen.Notch()
		fmt.Println("notch:", anchor, ok)
		if !ok {
			d := mygo.Screen.PrimaryDisplay()
			anchor = mygo.Rectangle{X: d.Bounds.X + d.Bounds.Width/2 - 100, Y: d.Bounds.Y, Width: 200, Height: 1}
		}
		fl := mygo.NewFlyout(mygo.FlyoutOptions{
			Screen:    true,
			Anchor:    anchor,
			Placement: mygo.PlacementBottom,
			Width:     max(anchor.Width+200, 400),
			Height:    200,
			Shadow:    true,
			Content: ui.View(func(c *ui.Context) {
				c.Root().Background(c.Theme().Surface)
				ui.Column(c).Fill().Center().Gap(8).Children(func() {
					ui.Textf(c, "notch: %v (has one: %v)", anchor, ok)
					ui.Text(c, "Press Escape to close")
				})
			}),
		})
		fl.OnClosed(func() { mygo.App.Quit() })
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
