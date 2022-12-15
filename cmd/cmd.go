package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New creates a new cobra command for a given year.
func New(year int, solutions map[int]func(string) error) *cobra.Command {
	cmd := Runner{
		year:      year,
		solutions: solutions,
	}

	return &cobra.Command{
		Use:    fmt.Sprintf("%d", year),
		Short:  fmt.Sprintf("Solutions for the Advent of Code %d.", year),
		Long:   fmt.Sprintf("Solutions for the Advent of Code %[1]d: https://adventofcode.com/%[1]d", year),
		PreRun: validateFlags,
		Run:    cmd.Run,
	}
}
func validateFlags(cmd *cobra.Command, _ []string) {
	file, _ := cmd.Flags().GetString("file")
	if file != "" {
		cmd.MarkFlagRequired("day")
	}
}

type Runner struct {
	year      int
	solutions map[int]func(string) error
}

func (c Runner) Run(cmd *cobra.Command, _ []string) {
	fmt.Printf("# Advent of Code %d\n", c.year)

	day, _ := cmd.Flags().GetInt("day")

	if day != 0 {
		c.runDay(cmd, day)
		return
	}

	fmt.Println("Day not specified, running everything!")
	c.runAll(cmd)

}

func (c Runner) runDay(cmd *cobra.Command, day int) {
	file, _ := cmd.Flags().GetString("file")
	test, _ := cmd.Flags().GetBool("test")

	if file == "" {
		file = defaultInputFile(c.year, day, test)
	}
	c.solve(day, file)
}

func (c Runner) runAll(cmd *cobra.Command) {
	test, _ := cmd.Flags().GetBool("test")
	for day := 1; day <= 25; day++ {
		file := defaultInputFile(c.year, day, test)
		c.solve(day, file)
	}
}

func defaultInputFile(year, day int, test bool) string {
	if test {
		return fmt.Sprintf("assets/inputs/%d/day%02d_test.txt", year, day)
	}

	return fmt.Sprintf("assets/inputs/%d/day%02d.txt", year, day)
}

func (c Runner) solve(day int, file string) {
	fn, implemented := c.solutions[day]
	if !implemented {
		fmt.Printf("- Day %02d not yet implemented\n", day)
		return
	}

	fmt.Println("\nDay", day)
	err := fn(file)
	if err != nil {
		fmt.Printf("Day %02d error: %v\n", day, err)
	}
}
