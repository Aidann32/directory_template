package os_utils

import (
	"errors"
	"regexp"
	"strings"
)

type WindowsUtils struct {
}

func (u WindowsUtils) ValidateFileName(name string) (err error) {
	if name == "" {
		return errors.New("file name cannot be empty")
	}

	invalidChars := `[\\/:*?"<>|]`
	match, err := regexp.MatchString(invalidChars, name)
	if err != nil {
		return err
	}
	if match {
		return errors.New("file name contains invalid characters")
	}

	invalidNames := []string{
		"CON", "PRN", "AUX", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}
	nameUpper := strings.ToUpper(name)
	for _, invalidName := range invalidNames {
		if nameUpper == invalidName || strings.HasPrefix(nameUpper, invalidName+".") {
			return errors.New("file name is reserved")
		}
	}

	if strings.HasSuffix(name, ".") || strings.HasSuffix(name, " ") {
		return errors.New("file name cannot end with a period or space")
	}

	if len(name) > 255 {
		return errors.New("file name cannot be longer than 255 characters")
	}

	return nil
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

func NewWindowsUtils() *WindowsUtils {
	return &WindowsUtils{}
}
