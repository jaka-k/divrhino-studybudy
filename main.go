/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/jaka-k/divrhino-studybudy/cmd"
	"github.com/jaka-k/divrhino-studybudy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
