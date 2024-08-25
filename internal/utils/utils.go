package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func DoesFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DoesDirectoryExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // Директория не существует
		}
		return false, err
	}

	if info.IsDir() {
		return true, nil // Директория существует
	}
	return false, nil
}

func IsGoInstalled() (err error) {
	if _, err := exec.Command("go", "version").Output(); err != nil {
		return errors.New("go is not installed on your machine")
	}
	return nil
}

func InitGoMod(dir, moduleName string) (err error) {
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	if _, err := os.Stat("go.mod"); err == nil {
		return fmt.Errorf("module already initialized")
	}

	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize module: %v", err)
	}

	return nil
}

func ParseProjectLayout(filePath string, projectLayout *map[string]interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	_ = json.Unmarshal(byteValue, &projectLayout)

	return nil
}

func ValidateModuleName(moduleName string) error {
	if strings.Contains(moduleName, " ") {
		return errors.New("module name contains spaces")
	}

	if strings.HasPrefix(moduleName, "/") || strings.HasSuffix(moduleName, "/") {
		return errors.New("module name should not start or end with a slash")
	}

	validPart := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

	parts := strings.Split(moduleName, "/")
	for _, part := range parts {
		if !validPart.MatchString(part) {
			return fmt.Errorf("invalid module name part: %s", part)
		}
		// Проверка на пустые части
		if len(part) == 0 {
			return errors.New("module name contains empty parts")
		}
	}

	return nil
}
