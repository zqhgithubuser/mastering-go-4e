/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// threeCmd represents the three command
var threeCmd = &cobra.Command{
	Use:     "three",
	Short:   "A brief description of your command",
	Aliases: []string{"cmd3"},
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("three called")
	},
}

func init() {
	rootCmd.AddCommand(threeCmd)
}
