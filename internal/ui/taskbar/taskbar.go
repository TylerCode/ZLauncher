// Package taskbar provides the bottom panel UI component for ZLauncher.
package taskbar

import (
	"image"
	"time"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/TylerCode/ZLauncher/internal/ui/theme"
)

// Height is the default taskbar height in dp.
const Height = 48

// Taskbar represents the bottom panel with start menu, running apps, and system tray.
type Taskbar struct {
	theme       *theme.ZorinTheme
	startButton widget.Clickable
}

// NewTaskbar creates a new taskbar with the given theme.
func NewTaskbar(th *theme.ZorinTheme) *Taskbar {
	return &Taskbar{
		theme: th,
	}
}

// Layout renders the taskbar.
func (t *Taskbar) Layout(gtx layout.Context) layout.Dimensions {
	// Set taskbar height
	gtx.Constraints.Min.Y = gtx.Dp(Height)
	gtx.Constraints.Max.Y = gtx.Dp(Height)

	// Draw taskbar background with slight transparency
	panelColor := t.theme.PanelDark
	panelColor.A = 0xE0 // Slight transparency for glass effect

	rect := image.Rectangle{Max: image.Point{X: gtx.Constraints.Max.X, Y: gtx.Constraints.Max.Y}}
	paint.FillShape(gtx.Ops, panelColor, clip.Rect(rect).Op())

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		// Start menu button (left)
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.layoutStartButton(gtx)
		}),
		// Running apps / dock area (center, flexible)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.layoutDockArea(gtx)
		}),
		// System tray with clock (right)
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.layoutSystemTray(gtx)
		}),
	)
}

// layoutStartButton renders the start menu button.
func (t *Taskbar) layoutStartButton(gtx layout.Context) layout.Dimensions {
	return layout.Inset{
		Left:  unit.Dp(8),
		Right: unit.Dp(8),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(t.theme.Theme, &t.startButton, "Z")
		btn.Background = t.theme.AccentColor
		btn.Color = t.theme.TextPrimary
		btn.CornerRadius = unit.Dp(4)
		return btn.Layout(gtx)
	})
}

// layoutDockArea renders the center dock/running apps area.
func (t *Taskbar) layoutDockArea(gtx layout.Context) layout.Dimensions {
	// Placeholder for app dock
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		label := material.Body2(t.theme.Theme, "")
		label.Color = t.theme.TextSecondary
		return label.Layout(gtx)
	})
}

// layoutSystemTray renders the system tray with clock.
func (t *Taskbar) layoutSystemTray(gtx layout.Context) layout.Dimensions {
	return layout.Inset{
		Left:  unit.Dp(8),
		Right: unit.Dp(12),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Display current time
		now := time.Now()
		timeStr := now.Format("3:04 PM")

		label := material.Body1(t.theme.Theme, timeStr)
		label.Color = t.theme.TextPrimary
		return label.Layout(gtx)
	})
}
