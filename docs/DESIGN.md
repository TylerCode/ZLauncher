# ZLauncher Design Document
## A Zorin OS-Inspired Launcher for Android Written in Go

---

## Table of Contents
1. [Executive Summary](#executive-summary)
2. [Feasibility Analysis](#feasibility-analysis)
3. [Design Goals](#design-goals)
4. [Zorin OS Design Philosophy](#zorin-os-design-philosophy)
5. [Architecture Overview](#architecture-overview)
6. [User Interface Design](#user-interface-design)
7. [Technical Implementation](#technical-implementation)
8. [Project Structure](#project-structure)
9. [Development Roadmap](#development-roadmap)
10. [Challenges and Mitigations](#challenges-and-mitigations)

---

## Executive Summary

ZLauncher is an Android launcher application inspired by the elegant and familiar desktop experience of Zorin OS. This document outlines the feasibility of developing such a launcher using Go, the design philosophy, architecture, and implementation strategy.

**Key Finding:** Yes, writing an Android launcher in Go is feasible. We will use [Gio](https://gioui.org/) as our UI framework for its GPU rendering performance and Android integration capabilities.

---

## Feasibility Analysis

### Can We Write an Android Launcher in Go?

**Yes**, with certain considerations and trade-offs.

### Available Approaches

#### 1. **Gio UI Framework** (Recommended)
- **Repository:** https://gioui.org/
- **Description:** Gio is a portable graphics library for Go that supports Android, iOS, Windows, macOS, Linux, and WebAssembly.
- **Pros:**
  - Pure Go implementation
  - Direct access to GPU rendering
  - Excellent performance
  - Active development and community
  - Native Android integration via JNI
- **Cons:**
  - Requires learning Gio's immediate-mode UI paradigm
  - Some Android-specific APIs need JNI bridges

#### 2. **Fyne Toolkit**
- **Repository:** https://fyne.io/
- **Description:** Cross-platform GUI toolkit following Material Design.
- **Pros:**
  - Easy to learn API
  - Good documentation
  - Material Design-like widgets
- **Cons:**
  - Less control over native Android features
  - May have larger binary size

#### 3. **gomobile (with custom UI)**
- **Repository:** https://github.com/golang/mobile
- **Description:** Go support for mobile platforms.
- **Pros:**
  - Direct Android NDK integration
  - Maximum control
- **Cons:**
  - More boilerplate required
  - Must implement UI from scratch

### Recommendation

**Gio UI is our chosen framework** because:
1. It provides the best balance of performance and flexibility
2. It allows fine-grained control over the UI for a custom Zorin-like experience
3. It has good Android integration for accessing launcher-specific APIs
4. It supports smooth animations and modern graphics

### Android Launcher Requirements

A launcher must implement these core Android concepts:
- `android.intent.category.HOME` - Register as home screen
- `android.intent.category.DEFAULT` - Default handler
- Package Manager queries for installed apps
- Widget hosting (AppWidgetHost)
- Wallpaper management

These will require JNI bridges from Go to Java/Kotlin.

---

## Design Goals

1. **Zorin OS Aesthetic** - Clean, modern, familiar desktop-like experience
2. **Performance** - Fast app launch times, smooth animations
3. **Customization** - Themes, layouts, and taskbar options
4. **Privacy-Focused** - No telemetry, no ads, open source
5. **Lightweight** - Minimal resource usage
6. **Accessibility** - Support for screen readers and accessibility features

---

## Zorin OS Design Philosophy

### Key Visual Elements to Adapt

1. **Taskbar/Panel**
   - Bottom-positioned panel (default) with option for top
   - Start menu button on the left
   - System tray on the right
   - Running apps in the center
   - Clock and date display

2. **Start Menu (App Drawer)**
   - Two-column layout
   - Left: Categories/Favorites
   - Right: Application list
   - Search bar at top
   - Power options at bottom

3. **Desktop**
   - Support for desktop icons
   - Wallpaper with blur effects
   - Right-click context menu

4. **Visual Style**
   - Rounded corners on elements
   - Subtle shadows and depth
   - Semi-transparent panels (glass effect)
   - Consistent icon sizing
   - Zorin Blue (#15A6F0) as accent color

### Zorin Layout Modes to Support

1. **Zorin Desktop (Default)**
   - Traditional desktop with bottom taskbar
   - Start menu in bottom-left

2. **Zorin Lite**
   - Similar to XFCE/traditional Windows
   - Two panels possible

3. **Zorin Touch**
   - Larger touch targets
   - Gesture-based navigation
   - Best for tablets

---

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                     ZLauncher Application                    │
├─────────────────────────────────────────────────────────────┤
│  UI Layer (Gio)                                             │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │   Taskbar   │ │ App Drawer  │ │   Desktop/Home      │   │
│  │  Component  │ │  Component  │ │    Component        │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
├─────────────────────────────────────────────────────────────┤
│  Service Layer (Go)                                         │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │    App      │ │  Settings   │ │     Theme           │   │
│  │   Manager   │ │   Manager   │ │    Manager          │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
├─────────────────────────────────────────────────────────────┤
│  Platform Bridge (JNI/cgo)                                  │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │  Package    │ │  Widget     │ │   Wallpaper         │   │
│  │  Manager    │ │   Host      │ │    Manager          │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
├─────────────────────────────────────────────────────────────┤
│  Android System                                             │
└─────────────────────────────────────────────────────────────┘
```

### Component Descriptions

#### UI Layer
- **Taskbar Component:** Renders the bottom panel with app shortcuts, system tray, and clock
- **App Drawer Component:** The start menu / app list overlay
- **Desktop Component:** Home screen with widget support and icon grid

#### Service Layer
- **App Manager:** Caches app list, handles app launches, manages favorites
- **Settings Manager:** Persists user preferences, layout configuration
- **Theme Manager:** Handles theming, accent colors, icon packs

#### Platform Bridge
- **Package Manager Bridge:** JNI calls to query installed applications
- **Widget Host Bridge:** Interface with Android's AppWidgetHost
- **Wallpaper Manager Bridge:** System wallpaper integration

---

## User Interface Design

### Screen Mockups (ASCII Art)

#### Default Home Screen
```
┌────────────────────────────────────────────────────┐
│                                                    │
│                    [Wallpaper]                     │
│                                                    │
│     📁          📱          🎵          ⚙️        │
│    Files       Phone       Music      Settings    │
│                                                    │
│     📷          💬          🌐          📧        │
│   Camera      Messages    Browser      Email      │
│                                                    │
│                                                    │
│                                                    │
├────────────────────────────────────────────────────┤
│ [Z]  📁 📱 🌐 📧  |  [  Running Apps  ]  | 🔋 📶 11:30│
└────────────────────────────────────────────────────┘
```

#### App Drawer (Start Menu)
```
┌────────────────────────────────────────────────────┐
│  [🔍 Search applications...                    ]  │
├─────────────────┬──────────────────────────────────┤
│                 │                                  │
│  ⭐ Favorites   │  Alarms              🕐         │
│                 │  Browser             🌐         │
│  📁 All Apps   │  Calculator          🔢         │
│                 │  Calendar            📅         │
│  🎮 Games      │  Camera              📷         │
│                 │  Clock               ⏰         │
│  📊 Office     │  Contacts            👥         │
│                 │  Email               📧         │
│  🛠️ System     │  Files               📁         │
│                 │  Gallery             🖼️         │
│                 │  Maps                🗺️         │
│                 │  Messages            💬         │
│                 │  Music               🎵         │
│                 │  Phone               📱         │
│                 │  Settings            ⚙️         │
├─────────────────┴──────────────────────────────────┤
│ 👤 User  |  ⚙️ Settings  |  🔒 Lock  |  ⏻ Power  │
└────────────────────────────────────────────────────┘
```

### Color Palette

| Color Name       | Hex Code  | Usage                    |
|------------------|-----------|--------------------------|
| Zorin Blue       | #15A6F0   | Accent, highlights       |
| Panel Dark       | #2D2D2D   | Taskbar background       |
| Panel Light      | #F5F5F5   | Light theme panels       |
| Text Primary     | #FFFFFF   | Primary text (dark mode) |
| Text Secondary   | #B0B0B0   | Secondary text           |
| Hover State      | #3D3D3D   | Button hover             |
| Active State     | #15A6F0   | Selected/active items    |

### Typography

- **Primary Font:** System default (Roboto on Android)
- **Title Size:** 18sp
- **Body Size:** 14sp
- **Caption Size:** 12sp

---

## Technical Implementation

### Go Module Structure

```go
module github.com/TylerCode/ZLauncher

go 1.21

require (
    gioui.org v0.4.1
    gioui.org/x v0.4.1
)
```

### Core Data Structures

```go
// App represents an installed application
type App struct {
    PackageName  string
    Label        string
    Icon         image.Image
    LaunchIntent string
    Category     string
    IsFavorite   bool
    LastUsed     time.Time
}

// Settings represents user preferences
type Settings struct {
    Theme           ThemeMode
    AccentColor     color.NRGBA
    TaskbarPosition Position
    IconSize        int
    ShowLabels      bool
    GridColumns     int
    GridRows        int
}

// ThemeMode represents the theme setting
type ThemeMode int

const (
    ThemeSystem ThemeMode = iota
    ThemeLight
    ThemeDark
)
```

### Key Interfaces

```go
// LauncherService defines the main service interface
type LauncherService interface {
    GetInstalledApps() ([]App, error)
    LaunchApp(packageName string) error
    GetAppIcon(packageName string) (image.Image, error)
    SetWallpaper(wallpaper image.Image) error
    GetWidgets() ([]Widget, error)
}

// ThemeService handles theming
type ThemeService interface {
    GetCurrentTheme() Theme
    SetTheme(theme Theme) error
    GetAccentColor() color.NRGBA
    SetAccentColor(c color.NRGBA) error
}

// StorageService handles persistence
type StorageService interface {
    SaveSettings(settings Settings) error
    LoadSettings() (Settings, error)
    SaveFavorites(apps []App) error
    LoadFavorites() ([]App, error)
}
```

### JNI Bridge Example

```go
//go:build android

package platform

/*
#include <jni.h>
#include <stdlib.h>

// Bridge functions declared in Java
extern jstring Java_com_zlauncher_Bridge_getInstalledApps(JNIEnv*, jobject);
extern void Java_com_zlauncher_Bridge_launchApp(JNIEnv*, jobject, jstring);
*/
import "C"

import (
    "encoding/json"
    "unsafe"
)

// GetInstalledApps retrieves all installed applications
func GetInstalledApps() ([]App, error) {
    // This would call into Java via JNI
    // Implementation details depend on Gio's Android integration
    return nil, nil
}

// LaunchApp starts an application by package name
func LaunchApp(packageName string) error {
    // JNI call to startActivity
    return nil
}
```

---

## Project Structure

```
ZLauncher/
├── cmd/
│   └── zlauncher/
│       └── main.go              # Application entry point
├── internal/
│   ├── ui/
│   │   ├── home/
│   │   │   ├── home.go          # Home screen component
│   │   │   └── grid.go          # App grid layout
│   │   ├── taskbar/
│   │   │   ├── taskbar.go       # Bottom panel
│   │   │   ├── startmenu.go     # Start menu/app drawer
│   │   │   └── systray.go       # System tray
│   │   ├── widgets/
│   │   │   └── widgets.go       # Widget support
│   │   └── theme/
│   │       ├── theme.go         # Theme definitions
│   │       └── colors.go        # Color palette
│   ├── service/
│   │   ├── apps.go              # App management
│   │   ├── settings.go          # Settings management
│   │   └── search.go            # App search functionality
│   └── platform/
│       ├── android/
│       │   ├── bridge.go        # JNI bridge
│       │   ├── apps.go          # Package manager queries
│       │   └── launcher.go      # Launcher registration
│       └── platform.go          # Platform interface
├── assets/
│   ├── icons/
│   │   └── zorin_logo.png       # ZLauncher logo
│   └── fonts/
│       └── ...                  # Custom fonts if needed
├── docs/
│   ├── DESIGN.md               # This document
│   └── API.md                  # API documentation
├── android/
│   ├── AndroidManifest.xml     # Android manifest
│   └── java/
│       └── com/zlauncher/
│           └── Bridge.java     # Java bridge code
├── go.mod
├── go.sum
├── Makefile                    # Build commands
└── README.md
```

---

## Development Roadmap

### Phase 1: Foundation (Weeks 1-4)
- [ ] Set up project structure with Gio
- [ ] Implement basic Android activity
- [ ] Create JNI bridge for package manager
- [ ] Basic app list retrieval
- [ ] Simple app launch functionality

### Phase 2: Core UI (Weeks 5-8)
- [ ] Implement taskbar component
- [ ] Create app drawer/start menu
- [ ] Home screen with app grid
- [ ] Basic theming (light/dark)
- [ ] App search functionality

### Phase 3: Zorin Polish (Weeks 9-12)
- [ ] Zorin-style visual design
- [ ] Animations and transitions
- [ ] Glass/blur effects on panels
- [ ] Icon pack support
- [ ] Favorites and categories

### Phase 4: Advanced Features (Weeks 13-16)
- [ ] Widget support
- [ ] Wallpaper management
- [ ] Gesture navigation
- [ ] Multiple layout modes
- [ ] Settings UI

### Phase 5: Polish & Release (Weeks 17-20)
- [ ] Performance optimization
- [ ] Accessibility features
- [ ] Beta testing
- [ ] Documentation
- [ ] F-Droid / Play Store submission

---

## Challenges and Mitigations

### Challenge 1: Android API Access
**Problem:** Many launcher features require Android-specific APIs not directly available in Go.

**Mitigation:** 
- Use JNI bridges through cgo
- Create a minimal Java/Kotlin layer for Android-specific functionality
- Leverage Gio's existing Android integration

### Challenge 2: Performance
**Problem:** Go garbage collection could cause UI stuttering.

**Mitigation:**
- Use object pooling for frequently allocated objects
- Profile and optimize hot paths
- Use Gio's efficient rendering pipeline
- Consider GOGC tuning for Android

### Challenge 3: Widget Hosting
**Problem:** AppWidgetHost is a complex Android API.

**Mitigation:**
- Defer widget support to later phases
- Start with basic widget support
- Consider using RemoteViews rendering from Java side

### Challenge 4: Icon Rendering
**Problem:** Android app icons are adaptive and can be complex.

**Mitigation:**
- Use Android's icon rendering and pass bitmaps to Go
- Support for adaptive icon backgrounds
- Cache rendered icons

### Challenge 5: Binary Size
**Problem:** Go binaries can be large.

**Mitigation:**
- Use `-ldflags="-s -w"` for smaller binaries
- Consider using TinyGo for smaller runtime (experimental)
- Strip debug symbols in release builds

---

## Conclusion

Building a Zorin OS-inspired launcher for Android in Go is **feasible and practical**. Using the Gio UI framework provides the best balance of performance, flexibility, and developer experience. While there are challenges with Android API access, these can be overcome with proper JNI bridging.

The resulting launcher would offer:
- A unique, desktop-like experience on Android
- Strong performance characteristics
- Privacy-focused, open-source design
- Familiar interface for desktop Linux users

This project represents an exciting opportunity to bring the elegant Zorin OS experience to Android users while demonstrating Go's capabilities for mobile development.

---

## References

- [Gio UI Framework](https://gioui.org/)
- [Fyne Toolkit](https://fyne.io/)
- [gomobile](https://github.com/golang/mobile)
- [Zorin OS](https://zorin.com/os/)
- [Android Launcher Development](https://developer.android.com/guide/topics/ui/look-and-feel/launchers)
- [Go Android Wiki](https://github.com/nicknogmail/goandroid/wiki)
