package os_utils

import (
	"errors"
	"regexp"
	"strings"
)

type WindowsUtils struct {
}

func (u WindowsUtils) InitGoMod(moduleName string) (err error) {
	return nil
}

func (u WindowsUtils) ValidateFileName(name string) (result bool) {

	return result
}

func (u WindowsUtils) ValidateDirectoryName(name string) (err error) {
	if name == "" {
		return errors.New("directory name cannot be empty")
	}

	invalidChars := `[\\/:*?"<>|]`
	match, err := regexp.MatchString(invalidChars, name)
	if err != nil {
		return err
	}
	if match {
		return errors.New("directory name contains invalid characters")
	}

	invalidNames := []string{
		"CON", "PRN", "AUX", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}
	for _, name := range invalidNames {
		if strings.EqualFold(name, name) {
			return errors.New("directory name is reserved")
		}
	}

	if strings.HasSuffix(name, ".") {
		return errors.New("directory name cannot end with a period")
	}

	if len(name) > 255 {
		return errors.New("directory name cannot be longer than 255 characters")
	}

	return nil

}

func (u WindowsUtils) DoesFileExist() (result bool) {
	return false
}

func (u WindowsUtils) IsGoInstalled() (err error) {

	return nil
}
