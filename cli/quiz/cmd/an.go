/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/pah-dev/fast-track-quiz/cli/quiz/utils"
	"github.com/spf13/cobra"
)

var anCmd = &cobra.Command{
	Use:   "an",
	Short: "Answer Question",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagId, _ := cmd.Flags().GetInt("id"); flagId != 0 {
			if flagAns, _ := cmd.Flags().GetInt("answer"); flagAns != 0 {
				services.AnswerQuestion(flagId, flagAns)
			}else{
				utils.PrintError("Answer number required, ex: -a 2")
			}
		}else{
			utils.PrintError("Question ID required, ex: -i 6")
		}
	},
}

func init() {
	anCmd.Flags().IntP("id", "i", 0, "Question ID")
	anCmd.Flags().IntP("answer", "a", 0, "Answer number")
	anCmd.MarkFlagRequired("id")
	anCmd.MarkFlagRequired("answer")
	rootCmd.AddCommand(anCmd)
}
