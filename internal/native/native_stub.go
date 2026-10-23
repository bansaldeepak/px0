//go:build !dev && !production

package native

import "net/http"

// Supported reports whether this binary was built with native window support.
func Supported() bool { return false }

// Run is a stub — rebuild with -tags wails to enable the native window.
func Run(_ http.Handler, _ string) error {
	panic("px0: not built with native window support; rebuild with: go build -tags wails .")
}
