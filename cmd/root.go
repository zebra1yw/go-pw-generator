/*
Copyright © 2023 zebra2 HERE github@zebra1yw
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwGen",
	Short: "A customizable password generator",
	Long: `pwGen is a powerful command-line tool designed to generate customizable passwords.
With pwGen, you can create passwords that are tailored to meet specific requirements. 

Example usage:
To generate a 12-character password with symbols, numbers, and mixed-case letters, simply type:

pwGen -l 12 -u -d -s`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}