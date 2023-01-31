package image

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvertPNGtoJPEG(path, OutputPath string, Quality int) error {
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

	fmt.Println("Converting file " + path + " to " + OutputPath + "/" + outputFile + ".jpeg")

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
