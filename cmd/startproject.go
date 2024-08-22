package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startprojectCmd represents the startproject command
var startprojectCmd = &cobra.Command{
	Use:   "startproject",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating project directory")
	},
}

func init() {
	rootCmd.AddCommand(startprojectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startprojectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startprojectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
