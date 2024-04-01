package main

import (
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"notepad/config"
	"notepad/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	createDataDir()
	// Create an instance of the app structure
	app := NewApp()

	docService := services.NewDocService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "notepad",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			docService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func createDataDir() {
	dir := config.Config.BaseDataPath
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				println(err.Error())
			}
		}
	}

}
