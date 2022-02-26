/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start Game",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flag, _ := cmd.Flags().GetString("name"); flag != "" {
			services.StartQuiz(flag)
		}
		// TODO: ver flag vacio
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "Guest", "Player Name")
	rootCmd.AddCommand(initCmd)
}

