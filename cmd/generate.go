/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/catusax/pgen/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := generator.ReadConfig()
		if err != nil {
			log.Fatalf("read config: %v", err)
		}

		g := getGenerator(cmd.Flags().GetString("engine"))

		registerFuncs(g)
		registerTemplates(g, conf.SoftFiles)

		if hard, err := cmd.Flags().GetBool("hard"); err == nil && hard {
			registerTemplates(g, conf.HardFiles)
		}

		bindings := generator.Bindings(conf.DefaultENVs).
			Set("NAME", args[0]).
			LoadFromFile().
			LoadFromENV()

		g.SetOptions(bindings)

		err = g.Generate()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	generateCmd.Flags().StringP("engine", "e", "text", "template engine liquid/text")
	generateCmd.Flags().BoolP("hard", "", false, "generate hardFiles defined in config file")
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
