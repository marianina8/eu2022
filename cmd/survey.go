/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// surveyCmd represents the survey command
var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "quick survey",
	Run: func(cmd *cobra.Command, args []string) {
		var questions = []*survey.Question{
			{
				Name: "prompt",
				Prompt: &survey.Input{
					Message: "What do you say to someone who wants to learn Go?",
				},
			},
		}
		answers := struct {
			Prompt string
		}{}
		_ = survey.Ask(questions, &answers)
		fmt.Println("You answered: ", answers.Prompt)
	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surveyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surveyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
