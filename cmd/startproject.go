package cmd

import (
	"fmt"
	"github.com/Aidann32/directory_template/internal/utils/os_utils"
	"github.com/spf13/cobra"
	"runtime"

	"github.com/Aidann32/directory_template/internal/service"
)

var startProjectCmd = &cobra.Command{
	Use:   "startproject [root_directory] [project_name] [module_name]",
	Short: "Starts project with specified layout",
	Long:  "If you do not pass your directory template, the standard one will be used",
	Run: func(cmd *cobra.Command, args []string) {
		var utils os_utils.OSUtils
		switch runtime.GOOS {
		case "windows":
			utils = os_utils.NewWindowsUtils()
		case "linux":
			utils = os_utils.NewUnixUtils()
		}

		if len(args) < 3 {
			fmt.Println("You must specify arguments")
		}

		layoutPath, _ := cmd.Flags().GetString("l")
		if layoutPath == "" {
			layoutPath = "project_layout.json"
		}

		if err := service.StartProject(args[0], args[1], args[2], layoutPath, utils); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Directory successfully created!")
	},
}

func init() {
	rootCmd.AddCommand(startProjectCmd)
	startProjectCmd.Flags().String("l", "", "Custom project layout")
}
