//go:build darwin

package darwin

import "github.com/ebitengine/purego/objc"

// notch: global pointer. NSEvent monitors of mouse moves, a global one
// for events of other apps and a local one for the app's own, which
// report the pointer's position (AppHandler.PointerMoved). AppKit gives
// global monitors the position of mouse moves without the Accessibility
// permission. Main thread only.
var pointerGlobal, pointerLocal id

func (a appController) WatchPointer(on bool) {
	if on == (pointerGlobal != 0) {
		return
	}
	if !on {
		send(class("NSEvent"), "removeMonitor:", uintptr(pointerGlobal))
		send(class("NSEvent"), "removeMonitor:", uintptr(pointerLocal))
		release(pointerGlobal)
		release(pointerLocal)
		pointerGlobal, pointerLocal = 0, 0
		return
	}
	const mask = 1<<5 | 1<<6 // mouse moved, left mouse dragged
	report := func() { a.b.h.PointerMoved(screen{}.CursorPoint()) }
	global := newBlock(func(_ objc.Block, ev id) { report() })
	local := newBlock(func(_ objc.Block, ev id) id { report(); return ev })
	pointerGlobal = retain(send(class("NSEvent"), "addGlobalMonitorForEventsMatchingMask:handler:", mask, uintptr(global)))
	pointerLocal = retain(send(class("NSEvent"), "addLocalMonitorForEventsMatchingMask:handler:", mask, uintptr(local)))
}
