package poweroftwo

import "testing"

func TestPowerOfTwo(t *testing.T) {
	powersOfTwo := []uint{2, 1024, 256, 512, 2048, 33554432}
	notPowersOfTwo := []uint{12, 624, 2156, 412, 148}

	for _, number := range powersOfTwo {
		if !isPowerOfTwo(number) {
			t.Errorf("Could not identify %d as power of 2\n", number)
		}
	}

	for _, number := range notPowersOfTwo {
		if isPowerOfTwo(number) {
			t.Errorf("%d is not a power of 2\n", number)
		}
	}

}
