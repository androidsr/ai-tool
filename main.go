package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.TempDir()
	}
	var dir = filepath.Join(configDir, "cache")
	os.Mkdir(dir, 0777)
	err = wails.Run(&options.App{
		Title:  "站点管理",
		Width:  1400,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               dir,
			WebviewBrowserPath:                "",
			Theme:                             windows.Light,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
