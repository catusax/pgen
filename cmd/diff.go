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
	Use:     "diff origin_file modified_file",
	Short:   "Print the diffmatchpatch result of two files",
	Long:    `Print the diffmatchpatch result of two files`,
	Example: "diff origin.txt modified.txt",
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

		diffs := dmp.DiffMain(string(file1), string(file2), false)

		fmt.Println(dmp.PatchToText(dmp.PatchMake(diffs)))
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
