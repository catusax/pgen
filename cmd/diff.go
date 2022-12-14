/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dmp := diffmatchpatch.New()

		if len(args) < 2 {
			panic("need two files")
		}

		file1, err := os.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		file2, err := os.ReadFile(args[1])
		if err != nil {
			panic(err)
		}

		// dmp.DiffMain(string(file1), string(file2), false)

		fmt.Println(dmp.DiffPrettyText(dmp.DiffMain(string(file1), string(file2), false))) // FIXME: print correct patch file https://neil.fraser.name/software/diff_match_patch/demos/patch.html
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
