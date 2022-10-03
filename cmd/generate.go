/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/catusax/pgen/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Regenerate current project",
	Long: `Regenerate current project

you can replace a file's template by putting a new template file in the
same directory. eg: [ main.go main.go.tmpl ] will use main.go.tmpl to
generate main.go.
	
view https://github.com/catusax/pgen for more information`,
	Run: func(cmd *cobra.Command, args []string) {
		g := getGenerator(cmd.Flags().GetString("engine"))

		registerFuncs(g)
		generator.LoadCustomFunction(g)

		registerTemplates(g, generator.Conf().SoftFiles)

		if hard, err := cmd.Flags().GetBool("hard"); err == nil && hard {
			registerTemplates(g, generator.Conf().HardFiles)
		}

		wd, _ := os.Getwd()

		bindings := generator.Bindings(generator.Conf().DefaultENVs).
			Set("NAME", filepath.Base(wd)).
			LoadFromFile().
			LoadFromENV()

		g.SetOptions(bindings)

		err := g.Generate()
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
