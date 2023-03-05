package cmd

import (
	"io"
	"os"
	"testing"

	"github.com/abdulahadakhter/png2jpeg/pkg/image"
)

func TestConvertPNGtoJPEG(t *testing.T) {
	path := "test.png"
	outputPath := "output"
	quality := 90
	expected := "Converting file test.png to output/test.jpeg\n"

	// Capture output of fmt.Println
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call function
	err := image.ConvertPNGtoJPEG(path, outputPath, quality)

	// Restore output
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(out) != expected {
		t.Errorf("Expected output %q, but got %q", expected, string(out))
	}
}
