// Package platform provides the interface for platform-specific functionality.
package platform

import "github.com/TylerCode/ZLauncher/internal/service"

// Platform defines the interface for platform-specific operations.
type Platform interface {
	// GetInstalledApps returns all installed applications.
	GetInstalledApps() ([]service.App, error)

	// LaunchApp launches an application by package name.
	LaunchApp(packageName string) error

	// SetAsDefaultLauncher prompts the user to set ZLauncher as default.
	SetAsDefaultLauncher() error

	// GetWallpaper returns the current system wallpaper.
	GetWallpaper() ([]byte, error)

	// SetWallpaper sets the system wallpaper.
	SetWallpaper(data []byte) error
}
