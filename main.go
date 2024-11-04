package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	imageprocessing "goroutines_pipeline/image_processing"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
				Image:     imageprocessing.ReadImage(p),
			}
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	// Prompt user to choose if they want to use Goroutines or not
	fmt.Println("Choose how to run the pipeline:")
	fmt.Println("1 - With Goroutines ")
	fmt.Println("2 - Without Goroutines")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var startTime time.Time
	var elapsedTime time.Duration

	if choice == "1" {
		fmt.Println("Running with concurrency...")
		startTime = time.Now()

		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}
		}

		elapsedTime = time.Since(startTime)
		fmt.Printf("Time taken with Goroutines: %v\n", elapsedTime)

	} else if choice == "2" {
		fmt.Println("Running without concurrency...")
		startTime = time.Now()

		for _, path := range imagePaths {
			// Sequentially read, process, and save each image
			checkInputFile(path)
			outPath := strings.Replace(path, "images/", "images/output/", 1)
			checkOutputPath(outPath)

			img := imageprocessing.ReadImage(path)
			img = imageprocessing.Resize(img)
			img = imageprocessing.Grayscale(img)
			imageprocessing.WriteImage(outPath, img)
			fmt.Println("Success!")
		}

		elapsedTime = time.Since(startTime)
		fmt.Printf("Time taken without Goroutines: %v\n", elapsedTime)

	} else {
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}

	fmt.Println("Press Enter to exit...")
	reader.ReadString('\n')
}

// Helper functions to check file and directory existence
func checkInputFile(inputPath string) {
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		log.Fatalf("Input image file %s does not exist: %v", inputPath, err)
	}
}

func checkOutputPath(outputPath string) {
	outputDir := filepath.Dir(outputPath)

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Unable to create output directory %s: %v", outputDir, err)
		}
	}
}
