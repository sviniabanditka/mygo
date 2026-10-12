//go:build windows && (amd64 || arm64)

package windows

import "github.com/egoist/mygo/internal/platform"

// notch: Screen.Notch. No display here has a notch.
func (screen) Notch() (platform.Rect, bool) { return platform.Rect{}, false }
