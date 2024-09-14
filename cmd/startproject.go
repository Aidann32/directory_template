package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Aidann32/directory_template/internal/static"
	"github.com/Aidann32/directory_template/internal/utils/os_utils"
	"github.com/spf13/cobra"
	"runtime"
	"strings"

	"github.com/Aidann32/directory_template/internal/service"
	projectUtils "github.com/Aidann32/directory_template/internal/utils"
)

var startProjectCmd = &cobra.Command{
	Use:   "startproject [root_directory] [project_name] [module_name]",
	Short: "Starts project with specified layout.",
	Long:  "If you do not pass your directory template, the standard one will be used.",
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
			return
		}

		var projectLayout map[string]interface{}
		layoutPath, _ := cmd.Flags().GetString("l")
		if len(strings.TrimSpace(layoutPath)) == 0 {
			_ = json.Unmarshal(static.DefaultLayout, &projectLayout)
		} else {
			if err := projectUtils.ParseProjectLayout(layoutPath, &projectLayout); err != nil {
				fmt.Printf("Error while reading project layout: %s", err.Error())
			}
		}

		if err := service.StartProject(args[0], args[1], args[2], projectLayout, utils); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Directory successfully created!")
	},
}

func init() {
	rootCmd.AddCommand(startProjectCmd)
	startProjectCmd.Flags().StringP("layout", "l", "", "Custom project layout")
}
