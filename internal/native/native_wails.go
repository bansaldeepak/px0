//go:build dev || production

package native

import (
	"context"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

// Supported reports whether this binary was built with native window support.
func Supported() bool { return true }

type app struct{ ctx context.Context }

func (a *app) startup(ctx context.Context) { a.ctx = ctx }

// Run opens a native desktop window using Wails, serving all requests through
// handler (the existing px0 HTTP server). The function blocks until the window
// is closed.
func Run(handler http.Handler, title string) error {
	a := &app{}
	return wails.Run(&options.App{
		Title:     title,
		Width:     1400,
		Height:    920,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Handler: handler,
		},
		OnStartup:        a.startup,
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 37, A: 255}, // Catppuccin Mocha base
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
	})
}
