//go:build android

// Package android provides Android-specific platform implementation.
package android

import (
	"github.com/TylerCode/ZLauncher/internal/platform"
	"github.com/TylerCode/ZLauncher/internal/service"
)

// AndroidPlatform implements the Platform interface for Android.
type AndroidPlatform struct {
	// TODO: Add JNI bridge references
}

// Ensure AndroidPlatform implements Platform interface.
var _ platform.Platform = (*AndroidPlatform)(nil)

// NewAndroidPlatform creates a new Android platform instance.
func NewAndroidPlatform() *AndroidPlatform {
	return &AndroidPlatform{}
}

// GetInstalledApps returns all installed applications.
func (p *AndroidPlatform) GetInstalledApps() ([]service.App, error) {
	// TODO: Implement JNI call to PackageManager
	// This will query Android's PackageManager for installed apps
	// with CATEGORY_LAUNCHER intent filter
	return []service.App{}, nil
}

// LaunchApp launches an application by package name.
func (p *AndroidPlatform) LaunchApp(packageName string) error {
	// TODO: Implement JNI call to start activity
	// Intent intent = packageManager.getLaunchIntentForPackage(packageName);
	// startActivity(intent);
	return nil
}

// SetAsDefaultLauncher prompts the user to set ZLauncher as default.
func (p *AndroidPlatform) SetAsDefaultLauncher() error {
	// TODO: Implement intent to show launcher picker
	return nil
}

// GetWallpaper returns the current system wallpaper.
func (p *AndroidPlatform) GetWallpaper() ([]byte, error) {
	// TODO: Implement JNI call to WallpaperManager
	return nil, nil
}

// SetWallpaper sets the system wallpaper.
func (p *AndroidPlatform) SetWallpaper(data []byte) error {
	// TODO: Implement JNI call to WallpaperManager
	return nil
}
