package imageprocessing

import (
	"image"
	"os"
	"testing"
)

func TestReadImage(t *testing.T) {
	img := ReadImage("test_images/test_image.jpeg")
	if img == nil {
		t.Errorf("Expected got image, got nothing")
	}
}

func TestResize(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	resizedImg := Resize(img)
	if resizedImg.Bounds().Dx() != 500 || resizedImg.Bounds().Dy() != 500 {
		t.Errorf("Expected image size 500x500, got %dx%d", resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy())
	}
}

func TestGrayscale(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	grayscaleImg := Grayscale(img)
	if grayscaleImg == nil {
		t.Errorf("Expected got grayscale image, got nothing")
	}
}

func TestWriteImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	outPath := "test_images/output_test_image.jpeg"
	WriteImage(outPath, img)

	_, err := os.Stat(outPath)
	if os.IsNotExist(err) {
		t.Errorf("Expected file at %s, but it does not exist", outPath)
	}

	os.Remove(outPath)
}
