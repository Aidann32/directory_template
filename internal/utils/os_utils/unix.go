package os_utils

import (
	"errors"
	"os/exec"
	"strings"
	"unicode/utf8"
)

type UnixUtils struct {
}

func (u UnixUtils) InitGoMod(moduleName string) (err error) {

	return nil
}

func (u UnixUtils) ValidateFileName(name string) (result bool) {

	return result
}

func (u UnixUtils) ValidateDirectoryName(name string) error {
	if name == "" {
		return errors.New("directory name cannot be empty")
	}

	if strings.Contains(name, "/") {
		return errors.New("directory name cannot contain '/'")
	}

	if strings.Contains(name, "\x00") {
		return errors.New("directory name cannot contain null character")
	}

	if utf8.RuneCountInString(name) > 255 {
		return errors.New("directory name cannot be longer than 255 characters")
	}
	return nil
}

func (u UnixUtils) DoesFileExist() (result bool) {
	return false
}

func (u UnixUtils) IsGoInstalled() (err error) {
	if _, err := exec.Command("go", "version").Output(); err != nil {
		return errors.New("go is not installed on your machine")
	}
	return nil
}
