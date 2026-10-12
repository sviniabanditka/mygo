// Notch-screen-flyout shows a flyout anchored to the top center of the
// primary display: a panel that hangs from the top edge of the screen,
// above the windows of other apps (FlyoutOptions.Screen).
//
//	go run ./examples/notch-screen-flyout
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

func main() {
	mygo.App.WhenReady(func() {
		d := mygo.Screen.PrimaryDisplay()
		anchor := mygo.Rectangle{X: d.Bounds.X + d.Bounds.Width/2 - 100, Y: d.Bounds.Y, Width: 200, Height: 1}
		fl := mygo.NewFlyout(mygo.FlyoutOptions{
			Screen:    true,
			Anchor:    anchor,
			Placement: mygo.PlacementBottom,
			Width:     480,
			Height:    240,
			Shadow:    true,
			Content: ui.View(func(c *ui.Context) {
				c.Root().Background(c.Theme().Surface)
				ui.Column(c).Fill().Center().Gap(8).Children(func() {
					ui.Text(c, "A screen flyout at the top center of the display").Bold()
					ui.Text(c, "Press Escape to close it")
				})
			}),
		})
		fl.OnClosed(func() { mygo.App.Quit() })
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
