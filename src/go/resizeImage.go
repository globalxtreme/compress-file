package main

import (
	"compress-file/support"
	"encoding/json"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"time"
)

func main() {
	start := time.Now()

	jsonData := os.Args[1]
	var params support.ParamImage
	if err := json.Unmarshal([]byte(jsonData), &params); err != nil {
		panic(err)
	}

	if len(params.Destination) == 0 {
		support.Error("Please set your destination file!!")
	} else {
		_, err := os.Stat(params.Destination)
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(params.Destination, 0777)
			} else {
				support.Error("Directory invalid!!")
			}
		}
	}

	quality := 70
	if params.Quality > 0 {
		quality = params.Quality
	}

	maxSize := params.MaxSize
	if maxSize == 0 {
		maxSize = 1000
	}

	file, err := os.Open(params.FilePath)
	if err != nil {
		support.Error(err.Error())
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		support.Error(err.Error())
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dx()

	var newWidth, newHeight int
	if width > height {
		newWidth = int(maxSize)
		newHeight = int(float64(maxSize) / float64(width) * float64(height))
	} else {
		newHeight = int(maxSize)
		newWidth = int(float64(maxSize) / float64(height) * float64(width))
	}

	img = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)

	out, err := os.Create(params.Destination + params.Filename)
	if err != nil {
		support.Error(err.Error())
	}
	defer out.Close()

	if err := webp.Encode(out, img, &webp.Options{Quality: float32(quality)}); err != nil {
		support.Error(err.Error())
	}

	fileInfo, err := out.Stat()
	if err != nil {
		support.Error(err.Error())
	}
	fileSize := fileInfo.Size()

	var message string
	if fileSize > int64(maxSize*1000) {
		message = fmt.Sprintf("The compressed image is larger than %d MB.", maxSize/1000)
	} else {
		elapsed := time.Since(start)
		message = fmt.Sprintf("The compressed image is smaller than %d MB. Execution time = %.3f", maxSize/1000, elapsed.Seconds())
	}

	support.Success(message, params.Destination+params.Filename)
}
