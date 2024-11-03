package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"os"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func LoadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			if _, err := os.Stat(p); err == nil {
				if _, err := os.Stat("images/output"); err == nil {
					job := Job{InputPath: p,
						OutPath: strings.Replace(p, "images/", "images/output/", 1)}
					job.Image = imageprocessing.ReadImage(p)
					out <- job
				}
			}

		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
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
		for job := range input { // Read from the channel
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
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	//arg := "other"

	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	if os.Args[1] == "concurrent" {
		channel1 := LoadImage(imagePaths)
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

	} else {

		for p := range imagePaths {
			path := []string{imagePaths[p]}
			channel1 := LoadImage(path)
			channel2 := resize(channel1)
			channel3 := convertToGrayscale(channel2)
			writeResults := saveImage(channel3)
			fmt.Print(writeResults)

		}
	}
}