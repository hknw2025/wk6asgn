
# Week 6 Assignment - Data Pipelines with Concurrency

This tools reads a series of images, resizes them, and converts them to grayscale. It can accomplish this either with or without concurrency using Go data pipelines.

## How to use this tool and examples:
* After cloning this git repository you can 
`go run main.go concurrent
* You will turn the normal images in the /images folder into grayscale resized images in the /images/Output folder. This will happen about a tenth of a second faster than 

`go run main.go other

because of the concurrent data pipelines. 

I was throughly impressed by how much faster using data pipelines are than using traditional coding methods. This is a very small data set that we are working with. We are doing some very simple operations to only 4 images. In my benchmark testing, the concurrent method was consistently about a tenth of a second faster than the non concurrent method. This is a large improvement in time considering that the whole operation occurred in less than a second and that these gains would compound greatly if we had more operations the images were undergoing or there were more images in the dataset. It really emphasizes how important of a tool concurrency is. 
















