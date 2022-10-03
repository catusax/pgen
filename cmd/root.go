/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pgen",
	Short: "A code generator for generating projects",
	Long: `pgen is a project generator that generates new projects based on your code template.
	
with this tool, you no longer need to copy-paste your project codes every time when you need to
create a new project. you cant wirte your own project skaffold with your favourite template engine,
then generate/regenerate a project with only one command.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var verbose int

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pgen.yaml)")

	rootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "print more logs, max level is -vvvv")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cobra.OnInitialize(func() {
		log.SetLevel(log.Level(verbose + 2)) // defaults to error level
	})
}
