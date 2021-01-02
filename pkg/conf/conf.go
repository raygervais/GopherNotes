package conf

import (
	"fmt"
	"os"
)

const (
	// LayoutISO acts a formatting helper when printing
	LayoutISO = "2006-01-02"
	// DatabaseLocation describes the default database name
	DatabaseLocation = "/notes.db"
	// ApplicationName descirbes the configuration folder name
	ApplicationName = "/gn"
	// DefaultEditor for *nix environments
	DefaultEditor = "vim"
)

// DetermineStorageLocation returns the results of UserConfigDir,
// which allows us to support per-os configurations being in the correct location.
func DetermineStorageLocation() (string, error) {
	return os.UserConfigDir()
}

// InitializeConfigurationLocation is a wrapper around the configuration folder
func InitializeConfigurationLocation(path string) error {
	appPath := path + ApplicationName

	if err := createConfigurationFolder(appPath); err != nil {
		return err
	}

	return nil
}

func createConfigurationFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0777); err != nil {
			return fmt.Errorf("Failed to create configuration folder: %s", err)
		}
	}

	return nil
}
