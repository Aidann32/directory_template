/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "embed"
	"encoding/json"
	"github.com/Aidann32/directory_template/cmd"
	"github.com/Aidann32/directory_template/internal/static"
)

// TODO: Parse layout when array used

/* TODO: New command "config" example: "config add key value --pd(project directory- project root directory where
config.json file is located) */

//go:embed assets/layouts/project_layout.json
var projectLayout []byte

//go:embed assets/file_contents/contents.json
var fileContents []byte

func main() {
	static.DefaultLayout = projectLayout
	_ = json.Unmarshal(fileContents, &static.FileContents)
	cmd.Execute()
}
