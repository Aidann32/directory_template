package os_utils

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type UnixUtils struct {
}

func (u UnixUtils) ValidateFileName(name string) (err error) {
	if name == "" {
		return errors.New("file name cannot be empty")
	}

	if strings.Contains(name, "/") {
		return errors.New("file name cannot contain '/'")
	}
	if strings.Contains(name, "\x00") {
		return errors.New("file name cannot contain null character")
	}

	if utf8.RuneCountInString(name) > 255 {
		return errors.New("file name cannot be longer than 255 characters")
	}

	return nil
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

func NewUnixUtils() *UnixUtils {
	return &UnixUtils{}
}
