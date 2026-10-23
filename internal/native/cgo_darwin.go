//go:build dev || production

package native

/*
#cgo LDFLAGS: -framework UniformTypeIdentifiers -framework AppKit
#include "dock_icon_darwin.h"
*/
import "C"
import "unsafe"

func setDockIcon(data []byte) {
	if len(data) == 0 {
		return
	}
	C.SetDockIcon(unsafe.Pointer(&data[0]), C.int(len(data)))
}
