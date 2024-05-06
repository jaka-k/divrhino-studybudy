/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note",
	Long:  `A Note can be anything.`,
}

func init() {
	rootCmd.AddCommand(noteCmd)

}
