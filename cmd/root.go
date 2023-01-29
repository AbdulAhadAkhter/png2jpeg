/*
Copyright © 2023 Abdul Ahad Akhter abdulahadakhter@outlook.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "png2jpeg",
	Short: "A small utility to convert png images to jpeg in current directory",
	Long:  `A small utility to convert png images to jpeg in current directory`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
