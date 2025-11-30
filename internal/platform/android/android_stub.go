//go:build !android

// Package android provides stubs for non-Android platforms (development/testing).
package android

import (
	"github.com/TylerCode/ZLauncher/internal/platform"
	"github.com/TylerCode/ZLauncher/internal/service"
)

// AndroidPlatform stub for non-Android platforms.
type AndroidPlatform struct{}

// Ensure AndroidPlatform implements Platform interface.
var _ platform.Platform = (*AndroidPlatform)(nil)

// NewAndroidPlatform creates a new Android platform stub.
func NewAndroidPlatform() *AndroidPlatform {
	return &AndroidPlatform{}
}

// GetInstalledApps returns mock apps for development.
func (p *AndroidPlatform) GetInstalledApps() ([]service.App, error) {
	// Return mock apps for desktop development
	return []service.App{
		{PackageName: "com.android.settings", Label: "Settings", Category: "System"},
		{PackageName: "com.android.chrome", Label: "Chrome", Category: "Internet"},
		{PackageName: "com.android.camera", Label: "Camera", Category: "Media"},
		{PackageName: "com.android.contacts", Label: "Contacts", Category: "Communication"},
		{PackageName: "com.android.phone", Label: "Phone", Category: "Communication"},
		{PackageName: "com.android.messaging", Label: "Messages", Category: "Communication"},
		{PackageName: "com.android.calculator", Label: "Calculator", Category: "Tools"},
		{PackageName: "com.android.calendar", Label: "Calendar", Category: "Tools"},
	}, nil
}

// LaunchApp is a stub for desktop development.
func (p *AndroidPlatform) LaunchApp(packageName string) error {
	return nil
}

// SetAsDefaultLauncher is a stub for desktop development.
func (p *AndroidPlatform) SetAsDefaultLauncher() error {
	return nil
}

// GetWallpaper is a stub for desktop development.
func (p *AndroidPlatform) GetWallpaper() ([]byte, error) {
	return nil, nil
}

// SetWallpaper is a stub for desktop development.
func (p *AndroidPlatform) SetWallpaper(data []byte) error {
	return nil
}
