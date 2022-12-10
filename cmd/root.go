package cmd

import (
	"os"

	puzzles2022 "github.com/leandrorondon/advent-of-code/pkg/2022"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Solutions for the Advent of Code puzzles.",
	Long:  `Solutions for the Advent of Code puzzles: https://adventofcode.com/`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().IntP("day", "d", 0, "Number of the day: 1 to 25")
	rootCmd.PersistentFlags().StringP("file", "f", "", "Input file (use only when specifying a day)")

	rootCmd.AddCommand(New(2022, puzzles2022.Solutions))
}
