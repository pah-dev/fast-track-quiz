/*
Copyright © 2022 JEAN PAUL BOBENRIETH <jeanpaulb79 (at) gmail com>

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
	Long: `
Answer question

quiz an -i X -a Y 

Where X is the question ID and Y is the answer number.
Ex: quiz an -i 5 -a 2`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagId, _ := cmd.Flags().GetInt("id"); flagId > 0 {
			if flagAns, _ := cmd.Flags().GetInt("answer"); flagAns != 0 {
				if flagAns > 0 && flagAns < 5 {
					services.AnswerQuestion(flagId, flagAns)
				}else{
					utils.PrintError("Answer must be between 1 and 4")
				}
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
