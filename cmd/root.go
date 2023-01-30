/*
Copyright © 2023 Abdul Ahad Akhter abdulahadakhter@outlook.com
*/
package cmd

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	InputPath  string
	OutputPath string
	Quality    int
)
var rootCmd = &cobra.Command{
	Use:   "png2jpeg",
	Short: "A small utility to convert png images to jpeg in an output directory",
	Long:  `A small utility to convert png images to jpeg in an output directory`,
	Run: func(cmd *cobra.Command, args []string) {

		if !strings.HasSuffix(InputPath, "/") {
			InputPath += "/"
		}
		paths, _ := filepath.Glob(InputPath + "*.png")
		for _, path := range paths {
			convertPNGtoJPEG(path, OutputPath, Quality)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&InputPath, "path", "p", "", "Path to images")
	rootCmd.Flags().StringVarP(&OutputPath, "output", "o", "", "Output path to convert images to")
	rootCmd.Flags().IntVarP(&Quality, "quality", "q", 0, "Output path to convert images to")
	rootCmd.MarkFlagRequired("path")
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}

func convertPNGtoJPEG(path, OutputPath string, Quality int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	if OutputPath == "" {
		OutputPath, _ = os.Getwd()
	}
	if _, err := os.Stat(OutputPath); os.IsNotExist(err) {
		os.Mkdir(OutputPath, os.ModePerm)
	}
	outputFile := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	fmt.Println("Converting file" + path + " to " + OutputPath + "/" + outputFile + ".jpeg")

	out, err := os.Create(OutputPath + "/" + outputFile + ".jpeg")
	if err != nil {
		return err
	}
	defer out.Close()

	if Quality == 0 {
		Quality = 80
	}

	return jpeg.Encode(out, img, &jpeg.Options{Quality: Quality})
}
