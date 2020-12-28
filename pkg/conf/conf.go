package conf

import (
	"fmt"
	"os"
)

const (
	// Formatting Helpers
	LayoutISO        = "2006-01-02"
	DatabaseLocation = "/notes.db"
	ApplicationName  = "/gn"
	DefaultEditor    = "vim"
)

func DetermineStorageLocation() (string, error) {
	return os.UserConfigDir()
}

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
