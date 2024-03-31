/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/vogonwann/studybudy/data"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new studybuddy note",
	Long:  `Creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

func init() {
	noteCmd.AddCommand(newCmd)
}

type PromptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red}}",
		Success: "{{ . | bold}}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	wordPromptContent := PromptContent{
		"Please provide a word",
		"What word would you like to make a note off",
	}

	word := promptGetInput(wordPromptContent)

	definitionPromptContent := PromptContent{
		"Please provide a definition.",
		fmt.Sprintf("What is the definition of %s?", word),
	}

	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := PromptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belongs to?", word),
	}
	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}

func promptGetSelect(pc PromptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
