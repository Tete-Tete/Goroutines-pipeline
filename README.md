# Image Processing Pipeline with Goroutines

This project is an image processing pipeline built in Go, based on a GitHub repository by Amrit Singh. The pipeline has been modified to add new features and enhancements. The pipeline reads images, resizes them, converts them to grayscale, and saves the modified images. The user can run the pipeline with or without Goroutines, allowing for a direct comparison of performance.

## Introduction
This project aims to demonstrate the benefits of concurrent programming by using Goroutines to process images in parallel. The pipeline can be run in two modes:
**Sequential Mode**: Processes images one by one.
**Concurrent Mode**: Uses Goroutines to process multiple images concurrently, taking advantage of Go's concurrency model and modern multi-core processors.

## Features
- **User Selection**: Choose between sequential and concurrent image processing.
- **Concurrency**: Demonstrates the power of Go's concurrency for speeding up image processing tasks.
- **Performance Benchmarking**: Compare performance with and without concurrency.
- **Error Handling**: Added error checks for file input and output to ensure stability.
- **Based on Amrit Singh's Work**: The project is built upon Amrit Singh's GitHub repository with additional modifications.

## Installation From Git and Setup
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/Goroutines-pipeline.git>
```

### Step 2: Run the Application
To build and run the Go application, run the following commands in your terminal:
```sh
go build -o pipeline.exe main.go

```
This will create an executable file named `pipeline` in your current directory. Then run the executable file, and you can choose use 1 or 2 and you will get different outputs which is the time for running this file. 

### Benchmark Results
- **Without Goroutines**: The benchmark results showed that the sequential execution was significantly slower.
- **With Goroutines**: The concurrent version effectively reduced the processing time, demonstrating the benefits of Go's concurrency model.

## Testing
### Running Tests
This project includes test cases to check the image_processing. To run tests, use:
```sh
go image_processing_test
```

## Conclusion
This project successfully demonstrates the power of Go's concurrency model in improving the performance of data pipelines. By comparing the sequential and concurrent versions of the image processing pipeline, it is evident that utilizing Goroutines can significantly reduce the processing time for large datasets. This highlights the efficiency and scalability that Go's concurrency can bring to real-world applications, particularly in data engineering and similar fields where throughput is critical.

