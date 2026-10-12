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

// notch: flyout key on click (platform.Flyout.KeyOnClick).

// takesKey reports whether a flyout may hold the keyboard: as it shows
// (Focusable) or once the user presses in it (KeyOnClick).
func takesKey(f platform.Flyout) bool { return f.Focusable || f.KeyOnClick }

// closeKeyOnClickFlyouts closes the visible KeyOnClick flyouts a press in
// the app lands outside of, as a press in another app closes them: the
// press's window is neither the flyout nor a window above it.
func (b *Backend) closeKeyOnClickFlyouts(ev id) {
	if send(b.app, "modalWindow") != 0 { // a dialog over the flyout (overFlyout)
		return
	}
	pressed := b.byNSWindow[send(ev, "window")]
	var outside []*window
	for _, w := range b.byNSWindow {
		if w.opts.Flyout == nil || w.popover != 0 || !w.flyout.KeyOnClick || w.closed || !sendBool(w.win, "isVisible") {
			continue
		}
		inside := false
		for k := pressed; k != nil; k = k.parent {
			if k == w {
				inside = true
				break
			}
		}
		if !inside {
			outside = append(outside, w)
		}
	}
	for _, w := range outside {
		if !w.closed && w.h.ShouldClose() {
			w.Close()
		}
	}
}

// notch: dialogs over flyouts.

// overFlyout prepares an application modal panel or alert whose parent is
// a flyout. A flyout does not show sheets: the sheet never appears and its
// caller waits for ever. runModal puts the dialog at the modal panel
// level, so the flyout goes down to the floating level while the dialog
// shows, else a flyout above the menu bar hides it. It returns the
// function that puts the flyout's level back.
func overFlyout(parent *window) func() {
	level := sendInt(parent.win, "level")
	send(parent.win, "setLevel:", nsFloatingWindowLevel)
	return func() {
		if !parent.closed {
			send(parent.win, "setLevel:", uintptr(level))
		}
	}
}
