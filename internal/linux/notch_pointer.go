//go:build linux && (amd64 || arm64)

package linux

// notch: global pointer. Not watched here.
func (appController) WatchPointer(bool) {}
