// Package home provides the home screen UI component for ZLauncher.
package home

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/TylerCode/ZLauncher/internal/ui/taskbar"
	"github.com/TylerCode/ZLauncher/internal/ui/theme"
)

// HomeScreen represents the main launcher home screen.
type HomeScreen struct {
	theme   *theme.ZorinTheme
	taskbar *taskbar.Taskbar
}

// NewHomeScreen creates a new home screen with the given theme.
func NewHomeScreen(th *theme.ZorinTheme) *HomeScreen {
	return &HomeScreen{
		theme:   th,
		taskbar: taskbar.NewTaskbar(th),
	}
}

// Layout renders the home screen.
func (h *HomeScreen) Layout(gtx layout.Context) layout.Dimensions {
	// Fill background
	fillBackground(gtx, h.theme.BackgroundColor)

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		// Main content area (desktop icons, widgets)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return h.layoutDesktop(gtx)
		}),
		// Bottom taskbar
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return h.taskbar.Layout(gtx)
		}),
	)
}

// layoutDesktop renders the desktop area with app icons.
func (h *HomeScreen) layoutDesktop(gtx layout.Context) layout.Dimensions {
	// Placeholder for desktop content
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		label := material.H4(h.theme.Theme, "ZLauncher")
		label.Color = h.theme.TextPrimary
		return layout.Inset{
			Top:    unit.Dp(100),
			Bottom: unit.Dp(100),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return label.Layout(gtx)
		})
	})
}

// fillBackground fills the entire context with the given color.
func fillBackground(gtx layout.Context, c color.NRGBA) {
	rect := image.Rectangle{Max: gtx.Constraints.Max}
	paint.FillShape(gtx.Ops, c, clip.Rect(rect).Op())
}
