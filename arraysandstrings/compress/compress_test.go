package compress

import "testing"

const (
	uncompressed = "aaaabbbdofhneiufhnefiuuuuyyyywwwa"
	compressed   = "a4b3dofhneiufhnefiu4y4w3a"
)

func TestCompress(t *testing.T) {

	compressedStr := compress(uncompressed)
	if compressed != compressedStr {
		t.Errorf("%s is not correct compression for %s\n", compressedStr, uncompressed)
	}

}
