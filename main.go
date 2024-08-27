/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "embed"
	"github.com/Aidann32/directory_template/cmd"
	"github.com/Aidann32/directory_template/internal/static"
)

//go:embed project_layout.json
var projectLayout []byte

func main() {
	static.DefaultLayout = projectLayout
	cmd.Execute()
}
