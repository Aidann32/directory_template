/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "embed"
	"github.com/Aidann32/directory_template/cmd"
	"github.com/Aidann32/directory_template/internal/static"
)

// TODO: Write to file  basic configuration if default layout is used

//go:embed assets/layouts/project_layout.json
var projectLayout []byte

//go:embed assets/file_contents/contents.json
var fileContents []byte

func main() {
	static.DefaultLayout = projectLayout
	static.FileContents = fileContents
	cmd.Execute()
}
