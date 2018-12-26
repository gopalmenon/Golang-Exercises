package rotateimage

import "testing"

const urlString = "http://farm1.static.flickr.com/122/263784734_c262172550.jpg"

func TestRotate(t *testing.T) {

	rotateImageBy90(urlString)

}
