package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	uppercase bool
	lowercase bool
	titlecase bool
	wordcount bool
	charcount bool
)

var rootCmd = &cobra.Command{
	Use:   "textfmt [text]",
	Short: "A simple text formatter",
	Long: `textfmt is a CLI tool for formatting and analyzing text.

You can provide text as arguments or pipe it in via stdin.
Multiple formatting options can be applied simultaneously.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi!")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "Convert text to uppercase")
	rootCmd.Flags().BoolVarP(&lowercase, "lower", "l", false, "Convert text to lowercase")
	rootCmd.Flags().BoolVarP(&titlecase, "title", "t", false, "Convert text to title case")
	rootCmd.Flags().BoolVar(&wordcount, "words", false, "Count words in text")
	rootCmd.Flags().BoolVar(&charcount, "chars", false, "Count characters in text")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
