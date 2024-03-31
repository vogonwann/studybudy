/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vogonwann/studybudy/data"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See a list of all notes you've added",
	Long:  `See a list of all notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
