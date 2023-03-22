package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version string

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Relingo Desktop",
		Width:  1200,
		Height: 800,
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "Relingo Desktop",
				Message: "\nVersion. " + version + "\n\nCopyright © 2022-2023 Bonaysoft.",
				Icon:    icon,
			},
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:  &options.RGBA{R: 250, G: 250, B: 250, A: 1},
		OnStartup:         app.startup,
		OnShutdown:        app.shutdown,
		HideWindowOnClose: true,
		Bind:              app.getBinds(),
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
