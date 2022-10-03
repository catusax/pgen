/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/catusax/pgen/generator"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new <project-name>",
	Short: "Generate a new project",
	Long: `Generate a new project

this command will generate a new project in new directory if it not exist

you must have a project template located in .template directory and 
a configuration file named .pgen_config.yaml. It will recursively search 
parent directories for a template and configuration file.

view https://github.com/catusax/pgen for more information
`,

	Example: "pgen new myproject",
	Run: func(cmd *cobra.Command, args []string) {
		g := getGenerator(cmd.Flags().GetString("engine"))

		registerFuncs(g)
		generator.LoadCustomFunction(g)
		registerTemplates(g, generator.Conf().OnceFiles, generator.Conf().HardFiles, generator.Conf().SoftFiles)

		bindings := generator.Bindings(generator.Conf().DefaultENVs).
			Set("NAME", args[0]).
			LoadFromENV()

		g.SetOptions(bindings)

		if _, err := os.Stat(args[0]); err == nil {
			fmt.Println("already exists!")
			os.Exit(1)
		}
		os.Mkdir(args[0], 0o755)
		os.Chdir(args[0])

		err := g.Generate()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func registerTemplates(g generator.Generator, fileGroups ...[]string) {
	for _, files := range fileGroups {
		for _, file := range files {
			err := g.Register(".template", file)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func registerFuncs(g generator.Generator) {
	g.RegisterFunc("dehyphen", func(s string) string {
		return strings.ReplaceAll(s, "-", "")
	})

	g.RegisterFunc("lower", strings.ToLower)

	g.RegisterFunc("title", func(s string) string {
		return strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, "_", "-")), "-", "")
	})

	g.RegisterFunc("dash", func(s string) string {
		return strings.ReplaceAll(s, "_", "-")
	})
}

func getGenerator(name string, err error) generator.Generator {
	if err != nil {
		return generator.NewTextGenerator()
	}
	switch name {
	case "liquid":
		return generator.NewLiquidGenerator()
	case "text":
		return generator.NewTextGenerator()
	default:
		return generator.NewTextGenerator()
	}
}

func init() {
	newCmd.Flags().StringP("engine", "e", "text", "template engine liquid/text")

	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
