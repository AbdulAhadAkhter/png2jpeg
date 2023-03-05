/*
Copyright © 2023 Abdul Ahad Akhter abdulahadakhter@outlook.com
*/
package cmd

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"os"

	"github.com/abdulahadakhter/png2jpeg/pkg/image"
	"github.com/spf13/cobra"
)

var (
	SingleFile  string
	BatchFolder string
	OutputPath  string
	Quality     int
)
var rootCmd = &cobra.Command{
	Use:    "png2jpeg",
	Short:  "A small utility to convert png images to jpeg in an output directory",
	Long:   `A small utility to convert png images to jpeg in an output directory`,
	PreRun: prerun,
	Run:    png2jpeg,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Init flags
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&SingleFile, "single", "s", "", "Path to images")
	rootCmd.Flags().StringVarP(&BatchFolder, "batch", "b", "", "Path to images")
	rootCmd.Flags().StringVarP(&OutputPath, "output", "o", "", "Output path to convert images to")
	rootCmd.Flags().IntVarP(&Quality, "quality", "q", 0, "Output path to convert images to")
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}

// Prerun script to ensure only one of single/batch flag is provided
func prerun(cmd *cobra.Command, args []string) {
	if SingleFile != "" && BatchFolder != "" {
		fmt.Println("Error: both variables provided, please provide only one")
		os.Exit(1)
	}
}

func png2jpeg(cmd *cobra.Command, args []string) {
	//If single file is provided and not empty then proceed
	if SingleFile != "" {
		//If single file provided is not a png then exit
		if !strings.HasSuffix(SingleFile, ".png") {
			fmt.Println("Can only convert png files")
			os.Exit(1)
		}
		//If single file provided does not exist then exit
		_, err := os.Stat(SingleFile)
		if os.IsNotExist(err) {
			fmt.Println("file does not exist:", SingleFile)
			os.Exit(1)
		}

		if OutputPath == "" {
			OutputPath = path.Dir(SingleFile) + "/output"
		}
		fmt.Println("Converting file: ", SingleFile)
		image.ConvertPNGtoJPEG(SingleFile, OutputPath, Quality)
	}
	//If batch folder provided and is not empty
	if BatchFolder != "" {
		//If batch folder provided and is not empty
		if strings.HasSuffix(BatchFolder, ".png") {
			fmt.Println("Only provide folders for tbe batch flag")
			return
		}
		//Handle case if the trailing forward slash isn't provided
		if !strings.HasSuffix(BatchFolder, "/") {
			BatchFolder += "/"
		}
		if OutputPath == "" {
			OutputPath = BatchFolder + "output"
		}
		//If folder does not exist
		_, err := os.Stat(BatchFolder)
		if os.IsNotExist(err) {
			fmt.Println("path does not exist:", BatchFolder)
			return
		}
		fmt.Println("Batch exporting files in: ", BatchFolder)
		paths, _ := filepath.Glob(BatchFolder + "*.png")
		for _, path := range paths {
			image.ConvertPNGtoJPEG(path, OutputPath, Quality)
		}
	}
}

// func png2jpeg(cmd *cobra.Command, args []string) {
// 	// Open the original image
// 	img, err := os.Open("original.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer img.Close()

// 	// Decode the image
// 	imgData, err := jpeg.Decode(img)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get the bounds of the original image
// 	bounds := imgData.Bounds()

// 	// Calculate the aspect ratio
// 	// aspectRatio := float64(bounds.Dx()) / float64(bounds.Dy())
// 	// fmt.Println("ASpect ratio", aspectRatio, "X dimensions", bounds.Dx(), "y dimensions", bounds.Dy())
// 	// Calculate the size of the border
// 	borderSize := 150
// 	newWidth := bounds.Dx() + borderSize*2
// 	newHeight := bounds.Dy() + borderSize*2
// 	bordercolor := "black"
// 	fmt.Println("Height", newHeight, "Width", newWidth)
// 	// Create a new image with equal borders
// 	var Border image.Uniform
// 	if bordercolor == "white" {
// 		Border = image.Uniform{color.RGBA{255, 255, 255, 255}}
// 	} else if bordercolor == "black" {
// 		Border = image.Uniform{color.RGBA{0, 0, 0, 255}}
// 	}

// 	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
// 	draw.Draw(newImg, newImg.Bounds(), &Border, image.Point{0, 0}, draw.Src)
// 	draw.Draw(newImg, bounds.Add(image.Point{borderSize, borderSize}), imgData, image.Point{0, 0}, draw.Src)

// 	// Create a new file to store the bordered image
// 	newFile, err := os.Create("bordered.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer newFile.Close()

// 	// Encode the bordered image as a JPEG and write it to the new file
// 	err = jpeg.Encode(newFile, newImg, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
