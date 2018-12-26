package rotateimage

import (
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
)

func rotateImageBy90(imageURL string) {

	response, err := http.Get(imageURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	imageObject, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	imageWidth, imageHeight := imageObject.Bounds().Max.X, imageObject.Bounds().Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, imageHeight, imageWidth))

	for x := 0; x < imageWidth; x += 1 {
		for y := 0; y < imageHeight; y += 1 {
			newImage.Set(y, x, imageObject.At(y, x))
		}
	}

	png.Encode(os.Stdout, newImage)

}
