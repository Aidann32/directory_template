package service

import (
	"fmt"
	"github.com/Aidann32/directory_template/internal/utils"
	"github.com/Aidann32/directory_template/internal/utils/os_utils"
	"os"
	"path/filepath"
	"strings"
)

func createLayout(layout map[string]interface{}, projectRoot string) error {
	for key, value := range layout {
		_ = os.Chdir(projectRoot)
		if strings.HasSuffix(key, "/") {
			if err := os.Mkdir(key, 0777); err != nil {
				return fmt.Errorf("error while creating layout: %s", err.Error())
			}
			_ = os.Chdir(key)
			switch value.(type) {
			case map[string]interface{}:
				dir, _ := os.Getwd()
				if err := createLayout(value.(map[string]interface{}), dir); err != nil {
					return err
				}
			case string:
				if value.(string) != "" {
					if _, err := os.Create(value.(string)); err != nil {
						return err
					}
				}
			}
		} else {
			if _, err := os.Create(key); err != nil {
				return err
			}
		}
	}
	return nil
}

func StartProject(rootDirectory, projectName, goModuleName, layoutPath string, osUtils os_utils.OSUtils) error {
	if err := utils.IsGoInstalled(); err != nil {
		return err
	}
	exists, err := utils.DoesDirectoryExist(rootDirectory)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("directory %s does not exist", rootDirectory)
	}

	if err := osUtils.ValidateDirectoryName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %s", err.Error())
	}

	if err := utils.ValidateModuleName(goModuleName); err != nil {
		return fmt.Errorf("invalid go module name: %s", err.Error())
	}

	var projectLayout map[string]interface{}
	if err := utils.ParseProjectLayout(layoutPath, &projectLayout); err != nil {
		return fmt.Errorf("error while reading project layout: %s", err.Error())
	}

	projectDirectory := filepath.Join(rootDirectory, projectName)
	if err := os.Mkdir(projectDirectory, 0777); err != nil {
		return fmt.Errorf("error while creating project root directory: %s", err.Error())
	}

	if err := os.Chdir(projectDirectory); err != nil {
		return fmt.Errorf("failed to change directory: %s", err)
	}

	if err := utils.InitGoMod(projectDirectory, goModuleName); err != nil {
		return fmt.Errorf("error while creating go module: %s", err.Error())
	}

	if err := createLayout(projectLayout, projectDirectory); err != nil {
		return fmt.Errorf("error while creating project layout: %s", err.Error())
	}

	return nil
}
