package config

import (
	"image/color"

	"golang.org/x/image/colornames"
)

type uiConfig struct {
	SquareSize int32

	WindowTitle string

	// colors
	BackgroundColor color.RGBA
	GridColor       color.RGBA

	TripDefaultColor    color.RGBA
	VehicleDefaultColor color.RGBA
}

type config struct {
	UI uiConfig
}

var Config config

func init() {
	Config.UI.SquareSize = 1
	Config.UI.BackgroundColor = colornames.Whitesmoke
	Config.UI.TripDefaultColor = colornames.Lightsteelblue
	Config.UI.VehicleDefaultColor = colornames.Green
	Config.UI.GridColor = colornames.Gray
	Config.UI.WindowTitle = "Google Hashcode 2018 - Simulator!"
}
