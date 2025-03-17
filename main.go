package main

import (
	"ExcelGeneratorWails/pkg/excel"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Criando uma nova instância do serviço Excel
	excelService := excel.NewExcelService()

	// Opções para o aplicativo Wails
	app := &options.App{
		Title:  "Gerador de Excel",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        excelService.Startup, // Inicializa o serviço
		Bind: []interface{}{
			excelService, // Expõe o serviço para o frontend
		},
	}

	// Iniciar o aplicativo
	if err := wails.Run(app); err != nil {
		log.Fatal(err)
	}
}