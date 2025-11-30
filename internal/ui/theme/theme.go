// Package theme provides the Zorin OS-inspired theming for ZLauncher.
package theme

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/widget/material"
)

// ZorinTheme holds the Zorin OS-inspired color palette and styling.
type ZorinTheme struct {
	*material.Theme

	// Zorin OS color palette
	AccentColor     color.NRGBA
	PanelDark       color.NRGBA
	PanelLight      color.NRGBA
	TextPrimary     color.NRGBA
	TextSecondary   color.NRGBA
	HoverState      color.NRGBA
	ActiveState     color.NRGBA
	BackgroundColor color.NRGBA
}

// Color constants from Zorin OS
var (
	ZorinBlue     = color.NRGBA{R: 0x15, G: 0xA6, B: 0xF0, A: 0xFF} // #15A6F0
	PanelDark     = color.NRGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF} // #2D2D2D
	PanelLight    = color.NRGBA{R: 0xF5, G: 0xF5, B: 0xF5, A: 0xFF} // #F5F5F5
	TextWhite     = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF} // #FFFFFF
	TextSecondary = color.NRGBA{R: 0xB0, G: 0xB0, B: 0xB0, A: 0xFF} // #B0B0B0
	HoverGray     = color.NRGBA{R: 0x3D, G: 0x3D, B: 0x3D, A: 0xFF} // #3D3D3D
	Transparent   = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
)

// NewZorinTheme creates a new Zorin OS-inspired theme.
func NewZorinTheme() *ZorinTheme {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection([]font.FontFace{}))
	th.Palette.Bg = PanelDark
	th.Palette.Fg = TextWhite
	th.Palette.ContrastBg = ZorinBlue
	th.Palette.ContrastFg = TextWhite

	return &ZorinTheme{
		Theme:           th,
		AccentColor:     ZorinBlue,
		PanelDark:       PanelDark,
		PanelLight:      PanelLight,
		TextPrimary:     TextWhite,
		TextSecondary:   TextSecondary,
		HoverState:      HoverGray,
		ActiveState:     ZorinBlue,
		BackgroundColor: PanelDark,
	}
}

// IsDarkMode returns true if the theme is in dark mode.
func (t *ZorinTheme) IsDarkMode() bool {
	return t.BackgroundColor == PanelDark
}

// SetDarkMode enables or disables dark mode.
func (t *ZorinTheme) SetDarkMode(dark bool) {
	if dark {
		t.BackgroundColor = PanelDark
		t.TextPrimary = TextWhite
		t.Palette.Bg = PanelDark
		t.Palette.Fg = TextWhite
	} else {
		t.BackgroundColor = PanelLight
		t.TextPrimary = PanelDark
		t.Palette.Bg = PanelLight
		t.Palette.Fg = PanelDark
	}
}
