/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ttycheckCmd represents the ttycheck command
var ttycheckCmd = &cobra.Command{
	Use:   "ttycheck",
	Short: "Checks if the terminal is an interactive tty",
	Long:  `Code to check if you’re outputting to a terminal…`,
	Run: func(cmd *cobra.Command, args []string) {
		fileInfo, _ := os.Stdout.Stat()
		if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("A terminal")
		} else {
			fmt.Println("Not a terminal")
		}

	},
}

func init() {
	rootCmd.AddCommand(ttycheckCmd)
}
