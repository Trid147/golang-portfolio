package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func extractLink(cmd *cobra.Command, args []string) {
	input := args[0]
	var url string

	markerIdx := strings.Index(input, "](")

	if markerIdx != -1 {
		urlStart := markerIdx + 2
		var urlEnd int

		closeBracket := strings.Index(input[urlStart:], ")")

		if closeBracket != -1 {
			urlEnd = urlStart + closeBracket
			url = input[urlStart:urlEnd]
		} else {
			url = input
		}
	} else {
		url = input
	}

	fmt.Printf("Result: %s\n", url)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "link-extractor [TEXT_OR_URL]",
		Short: "A Markdown-to-HTML Link Extractor",
		Long:  "A CLI utility that extracts raw URLs from Markdown formatted links like [Name](url).",
		Args:  cobra.ExactArgs(1),
		Run:   extractLink,
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
