package service

import (
	"fmt"
	"github.com/Aidann32/directory_template/internal/utils/os_utils"
	"os"
)

func StartProject(rootDirectory, projectName string, os_utils os_utils.OSUtils) error {
	if err := os_utils.IsGoInstalled(); err != nil {
		return err
	}
	if err := os_utils.ValidateDirectoryName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %s", err.Error())
	}

	err := os.Mkdir(projectName, 0777)
	if err != nil {
		return fmt.Errorf("error while creating project directory: %s", err.Error())
	}

	return nil
}
