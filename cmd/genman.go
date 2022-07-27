/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// genmanCmd represents the command to generate a man page
var genmanCmd = &cobra.Command{
	Use:   "genman",
	Short: "generate man page",
	Long:  `this command will generate a man pages for the root command in the /tmp dir`,
	Run: func(cmd *cobra.Command, args []string) {
		header := &doc.GenManHeader{
			Title:   "Application Man Pages",
			Section: "1",
		}
		_ = doc.GenManTree(rootCmd, header, "/tmp")

	},
}

func init() {
	rootCmd.AddCommand(genmanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genmanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genmanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
