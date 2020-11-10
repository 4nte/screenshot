package main

import (
	"fmt"
	"github.com/4nte/screenshot"
	"image"
	"image/png"
	"os"
	"time"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func takeScreen() {
	// Capture each displays.

	xgbConn, err := screenshot.NewXgbConnection()
	if err != nil {
		panic(err)
	}

	n := xgbConn.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := xgbConn.GetDisplayBounds(i)
		all = bounds.Union(all)

		_, err := xgbConn.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		//fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		//save(img, fileName)

		//fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	// Capture all desktop region into an image.
	fmt.Printf("%v\n", all)
	_, err = xgbConn.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	//save(img, "all.png")
}

func main() {

	for true {
		takeScreen()
		<-time.After(1 * time.Millisecond)
	}
}
