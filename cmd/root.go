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
	SingleFile  string
	BatchFolder string
	OutputPath  string
	Quality     int
)
var rootCmd = &cobra.Command{
	Use:   "png2jpeg",
	Short: "A small utility to convert png images to jpeg in an output directory",
	Long:  `A small utility to convert png images to jpeg in an output directory`,
	// PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {

		if SingleFile != "" && BatchFolder != "" {
			fmt.Println("Error: both variables provided, please provide only one")
			return
		}
		if SingleFile != "" {

			if !strings.HasSuffix(SingleFile, ".png") {
				fmt.Println("Can only convert png files")
				return
			}

			_, err := os.Stat(SingleFile)
			if os.IsNotExist(err) {
				fmt.Println("file does not exist:", SingleFile)
				return
			}

			fmt.Println("Converting file: ", SingleFile)
			convertPNGtoJPEG(SingleFile, OutputPath, Quality)
		}
		if BatchFolder != "" {
			if strings.HasSuffix(BatchFolder, ".png") {
				fmt.Println("Only provide folders for tbe batch flag")
				return
			}
			if !strings.HasSuffix(BatchFolder, "/") {
				BatchFolder += "/"
			}
			_, err := os.Stat(BatchFolder)
			if os.IsNotExist(err) {
				fmt.Println("path does not exist:", BatchFolder)
				return
			}
			fmt.Println("Batch exporting files in: ", BatchFolder)
			paths, _ := filepath.Glob(BatchFolder + "*.png")
			for _, path := range paths {
				convertPNGtoJPEG(path, OutputPath, Quality)
			}
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
	rootCmd.Flags().StringVarP(&SingleFile, "single", "s", "", "Path to images")
	rootCmd.Flags().StringVarP(&BatchFolder, "batch", "b", "", "Path to images")
	rootCmd.Flags().StringVarP(&OutputPath, "output", "o", "", "Output path to convert images to")
	rootCmd.Flags().IntVarP(&Quality, "quality", "q", 0, "Output path to convert images to")
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
