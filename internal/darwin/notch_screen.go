//go:build darwin

package darwin

import "github.com/egoist/mygo/internal/platform"

// notch: Screen.Notch. The camera housing of a built-in display, from
// NSScreen's safe area (macOS 12 and later): the gap between the auxiliary
// top-left and top-right areas, as tall as the top inset.
func (screen) Notch() (platform.Rect, bool) {
	var r platform.Rect
	var ok bool
	withPool(func() {
		for _, s := range arrayItems(send(class("NSScreen"), "screens")) {
			if !respondsTo(s, "safeAreaInsets") {
				return
			}
			// NSEdgeInsets{top, left, bottom, right} has the layout of
			// NSRect{{x, y}, {w, h}}: the same four doubles.
			insets := msgRect(s, sel("safeAreaInsets"))
			top := insets.Origin.X
			if top <= 0 {
				continue
			}
			frame := msgRect(s, sel("frame"))
			left := msgRect(s, sel("auxiliaryTopLeftArea"))
			right := msgRect(s, sel("auxiliaryTopRightArea"))
			x0, x1 := left.Origin.X+left.Size.Width, right.Origin.X
			r = rectFromMac(NSRect{
				Origin: NSPoint{x0, frame.Origin.Y + frame.Size.Height - top},
				Size:   NSSize{x1 - x0, top},
			})
			ok = true
			return
		}
	})
	return r, ok
}
