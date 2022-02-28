/*
Copyright © 2022 JEAN PAUL BOBENRIETH <jeanpaulb79 (at) gmail com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Super Simple Quiz",
	Long: `
Super Simple Quiz

A simple quiz with a few questions and a few alternatives for each question.`,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configFile = ".quiz-config.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading configuration file: ", viper.ConfigFileUsed())
	}
}

