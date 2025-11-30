// Package service provides business logic services for ZLauncher.
package service

import (
	"image"
	"time"
)

// App represents an installed application.
type App struct {
	PackageName  string
	Label        string
	Icon         image.Image
	LaunchIntent string
	Category     string
	IsFavorite   bool
	LastUsed     time.Time
}

// AppService manages installed applications.
type AppService struct {
	apps      []App
	favorites []App
}

// NewAppService creates a new app service.
func NewAppService() *AppService {
	return &AppService{
		apps:      make([]App, 0),
		favorites: make([]App, 0),
	}
}

// GetApps returns all installed applications.
func (s *AppService) GetApps() []App {
	return s.apps
}

// GetFavorites returns favorite applications.
func (s *AppService) GetFavorites() []App {
	return s.favorites
}

// SetFavorite marks an app as favorite.
func (s *AppService) SetFavorite(packageName string, favorite bool) {
	for i := range s.apps {
		if s.apps[i].PackageName == packageName {
			s.apps[i].IsFavorite = favorite
			if favorite {
				s.favorites = append(s.favorites, s.apps[i])
			} else {
				// Remove from favorites
				for j := range s.favorites {
					if s.favorites[j].PackageName == packageName {
						s.favorites = append(s.favorites[:j], s.favorites[j+1:]...)
						break
					}
				}
			}
			break
		}
	}
}

// RefreshApps reloads the app list from the platform.
// This is a placeholder - actual implementation will use platform bridge.
func (s *AppService) RefreshApps() error {
	// TODO: Implement platform-specific app loading
	return nil
}

// LaunchApp launches an application by package name.
// This is a placeholder - actual implementation will use platform bridge.
func (s *AppService) LaunchApp(packageName string) error {
	// TODO: Implement platform-specific app launching
	for i := range s.apps {
		if s.apps[i].PackageName == packageName {
			s.apps[i].LastUsed = time.Now()
			break
		}
	}
	return nil
}
