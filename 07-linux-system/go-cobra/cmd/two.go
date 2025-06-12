/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twoCmd represents the two command
var twoCmd = &cobra.Command{
	Use:     "two",
	Short:   "A brief description of your command",
	Aliases: []string{"cmd2"},
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("two called")
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)

	twoCmd.Flags().StringP("username", "u", "Mike", "Username")
}
