package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"os"
	"strings"
	"testing"
)

func TestLoadImage(t *testing.T) {

	ans := Job{
		InputPath: "images/image1.jpeg",
		OutPath:   "images/output/image1.jpeg",
	}

	paths := []string{
		"images/image1.jpeg",
	}

	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			if _, err := os.Stat(p); err == nil {
				if _, err := os.Stat("images/output"); err == nil {
					job := Job{InputPath: p,
						OutPath: strings.Replace(p, "images/", "images/output/", 1)}
					// for some reason these were a bit off when they being read in
					job.Image = imageprocessing.ReadImage(p)
					ans.Image = imageprocessing.ReadImage(p)
					out <- job
				}
			}

		}
		close(out)
	}()

	// compare parsed data to the data it should be parsed to
	test := <-out
	if test.InputPath != ans.InputPath {
		t.Errorf("parsed incorrectly")
	}
	if test.OutPath != ans.OutPath {
		t.Errorf("parsed incorrectly")
	}

}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
