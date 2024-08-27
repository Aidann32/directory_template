/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "embed"
	"github.com/Aidann32/directory_template/cmd"
)

//go:embed project_layout.json
var projectLayout []byte

func main() {
	cmd.Execute(projectLayout)
}
