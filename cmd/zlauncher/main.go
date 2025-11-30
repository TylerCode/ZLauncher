// Package main is the entry point for ZLauncher - a Zorin OS-inspired Android launcher.
package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"github.com/TylerCode/ZLauncher/internal/ui/home"
	"github.com/TylerCode/ZLauncher/internal/ui/theme"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("ZLauncher"))

		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	th := theme.NewZorinTheme()
	homeScreen := home.NewHomeScreen(th)

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			homeScreen.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
