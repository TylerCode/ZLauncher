# ZLauncher

A Zorin OS-inspired launcher for Android, written in Go using the [Gio](https://gioui.org/) UI framework.

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

- **Language:** Go 1.21+
- **UI Framework:** [Gio](https://gioui.org/)
- **Platform:** Android (via Gio Android support)

## Building

### Prerequisites

For desktop development/testing:
```bash
# Ubuntu/Debian
sudo apt-get install libxkbcommon-dev libwayland-dev libgl1-mesa-dev \
    libvulkan-dev libxkbcommon-x11-dev libx11-xcb-dev libxcursor-dev \
    libxfixes-dev libegl1-mesa-dev xorg-dev
```

### Build Commands

```bash
# Build for desktop (development)
make build

# Run desktop version
make run

# Build Android APK (requires gogio)
make apk

# Run tests
make test

# Lint code
make lint
```

### Install gogio for Android builds

```bash
go install gioui.org/cmd/gogio@latest
```

## Project Status

🚀 **v1 Development** - Basic project structure with Gio framework in place.

### Current Progress
- [x] Project structure setup
- [x] Gio framework integration
- [x] Zorin OS theme/color palette
- [x] Basic taskbar component
- [x] Home screen scaffold
- [ ] App drawer/start menu
- [ ] Desktop icon grid
- [ ] Android platform bridge
- [ ] App launching functionality

## Contributing

Contributions are welcome! Please read the design document first to understand the project architecture and goals.

## License

[To be determined]

## Acknowledgments

- [Zorin OS](https://zorin.com/os/) for the design inspiration
- [Gio UI](https://gioui.org/) for making Go GUI development possible on mobile