package unsupported

import "github.com/egoist/mygo/internal/platform"

// notch: Screen.Notch. No display here has a notch.
func (screen) Notch() (platform.Rect, bool) { return platform.Rect{}, false }
