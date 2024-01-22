// Package main contains a main application code for GUI game
package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"

	"github.com/AllenDang/giu"

	"github.com/gucio321/tic-tac-go/pkg/giuwidget"
)

//go:embed logo.png
var logoData []byte

func main() {
	const (
		screenX, screenY = 640, 480
	)

	logo, err := png.Decode(bytes.NewReader(logoData))
	if err != nil {
		log.Fatal("error decoding logo bytes")
	}

	wnd := giu.NewMasterWindow("Tic-Tac-Go", screenX, screenY, 0)
	wnd.SetIcon([]image.Image{logo})
	wnd.Run(func() {
		giu.SingleWindow().Layout(giuwidget.Game())
	})
}
