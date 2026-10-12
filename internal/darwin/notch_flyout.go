//go:build darwin

package darwin

import "github.com/egoist/mygo/internal/platform"

// notch: screen flyout. A flyout anchored to a rectangle of the screen
// (platform.Flyout.Screen): it has no parent and no tray, floats above
// the windows of other apps, shows on every Space and over full screen
// apps, and is placed in the bounds of its display, the menu bar
// included.

const collectionCanJoinAllSpaces = 1 << 0

// setupScreenFlyout gives a new screen flyout its level and spaces.
func (w *window) setupScreenFlyout() {
	send(w.win, "setLevel:", nsFloatingWindowLevel)
	send(w.win, "setCollectionBehavior:", collectionCanJoinAllSpaces|collectionTransient|collectionIgnoresCycle|collectionFullScreenAuxiliary)
}

// placeScreenFlyout places a screen flyout against its anchor in the
// bounds of the display the anchor is on, or the primary display's.
func (w *window) placeScreenFlyout(f platform.Flyout, size platform.Size, displays []platform.Display) {
	bounds := displayBoundsFor(f.Anchor, displays)
	r, _, _ := f.Place(f.Anchor, size, bounds)
	msgSetRectBool(w.win, sel("setFrame:display:"), rectToMac(r), true)
}

// displayBoundsFor returns the bounds of the display that holds the
// center of r, or of the primary display.
func displayBoundsFor(r platform.Rect, displays []platform.Display) platform.Rect {
	cx, cy := r.X+r.Width/2, r.Y+r.Height/2
	for _, d := range displays {
		b := d.Bounds
		if cx >= b.X && cx < b.X+b.Width && cy >= b.Y && cy < b.Y+b.Height {
			return b
		}
	}
	for _, d := range displays {
		if d.Primary {
			return d.Bounds
		}
	}
	if len(displays) > 0 {
		return displays[0].Bounds
	}
	return platform.Rect{}
}
