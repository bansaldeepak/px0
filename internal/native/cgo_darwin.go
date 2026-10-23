//go:build dev || production

package native

// #cgo LDFLAGS: -framework UniformTypeIdentifiers
import "C"
