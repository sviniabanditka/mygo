//go:build windows && (amd64 || arm64)

package windows

// notch: global pointer. Not watched here.
func (appController) WatchPointer(bool) {}
