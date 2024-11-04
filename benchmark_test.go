package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"testing"
)

// Benchmark for running the pipeline without Goroutines
func BenchmarkPipelineWithoutGoroutines(b *testing.B) {
	imagePaths := []string{
		"test_images/image1.jpeg",
		"test_images/image2.jpeg",
		"test_images/image3.jpeg",
		"test_images/image4.jpeg",
	}

	for i := 0; i < b.N; i++ {
		for _, path := range imagePaths {
			img := imageprocessing.ReadImage(path)
			img = imageprocessing.Resize(img)
			img = imageprocessing.Grayscale(img)
			imageprocessing.WriteImage("test_images/output/"+path, img)
		}
	}
}

// Benchmark for running the pipeline with Goroutines
func BenchmarkPipelineWithGoroutines(b *testing.B) {
	imagePaths := []string{
		"test_images/image1.jpeg",
		"test_images/image2.jpeg",
		"test_images/image3.jpeg",
		"test_images/image4.jpeg",
	}

	for i := 0; i < b.N; i++ {
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		// Wait for all images to be saved
		for range writeResults {
		}
	}
}
