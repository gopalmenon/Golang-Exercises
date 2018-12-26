package rotation

import "testing"

const (
	original = "waterbottle"
	rotated  = "terbottlewa"
)

func TestRotation(t *testing.T) {

	if !IsRotationOf(original, rotated) {
		t.Errorf("Could not identify %s as rotation of %s\n", rotated, original)
	}

}
