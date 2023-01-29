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
	FilePath   string
	OutputPath string
)
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert png images to jpeg",
	Long:  `Convert png images to jpeg`,
	Run: func(cmd *cobra.Command, args []string) {
		paths, _ := filepath.Glob(FilePath + "*.png")
		for _, path := range paths {
			convertPNGtoJPEG(path, OutputPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&FilePath, "path", "p", "", "Path to images")
	convertCmd.Flags().StringVarP(&OutputPath, "output", "o", "", "Output path to convert images to")
}

func convertPNGtoJPEG(path, OutputPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
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

	return jpeg.Encode(out, img, nil)
}
