# ZLauncher

A Zorin OS-inspired launcher for Android, written in Go.

## Overview

ZLauncher brings the elegant, familiar desktop experience of [Zorin OS](https://zorin.com/os/) to Android devices. It features a traditional desktop-style interface with a taskbar, start menu, and desktop icons - perfect for users who prefer a desktop-like experience on their mobile devices or tablets.

## Key Features (Planned)

- 🖥️ **Desktop-Style Interface** - Bottom taskbar with start menu, system tray, and clock
- 🎨 **Zorin OS Aesthetics** - Clean, modern design with Zorin Blue accents
- 🔍 **App Search** - Quick search through all installed applications
- ⭐ **Favorites & Categories** - Organize apps your way
- 🌙 **Dark/Light Themes** - Match your system theme or choose your own
- 🔒 **Privacy-Focused** - No telemetry, no ads, open source

## Documentation

- [Design Document](docs/DESIGN.md) - Comprehensive design specification including feasibility analysis, architecture, and implementation details

## Tech Stack

- **Language:** Go
- **UI Framework:** [Gio](https://gioui.org/) (recommended) or [Fyne](https://fyne.io/)
- **Platform:** Android (via gomobile/Gio Android support)

## Is Writing an Android Launcher in Go Possible?

**Yes!** See our [Design Document](docs/DESIGN.md#feasibility-analysis) for a detailed feasibility analysis. The key enabler is the [Gio UI framework](https://gioui.org/), which provides cross-platform GUI capabilities including Android support with GPU-accelerated rendering.

## Project Status

🚧 **Planning Phase** - Currently in the design and planning stage.

## Contributing

Contributions are welcome! Please read the design document first to understand the project architecture and goals.

## License

[To be determined]

## Acknowledgments

- [Zorin OS](https://zorin.com/os/) for the design inspiration
- [Gio UI](https://gioui.org/) for making Go GUI development possible on mobile