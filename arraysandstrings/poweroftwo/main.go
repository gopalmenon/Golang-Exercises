package poweroftwo

func isPowerOfTwo(number uint) bool {

	return number != 0 && number&(number-1) == 0

}
