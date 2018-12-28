package towerofhanoi

import "fmt"

func towerOfHanoi(numberOfDisks, fromPost, toPost, bufferPost int) {

	if numberOfDisks == 0 {
		return
	}

	towerOfHanoi(numberOfDisks-1, fromPost, bufferPost, toPost)

	fmt.Printf("Move a disk from post %d to post %d\n", fromPost, toPost)

	towerOfHanoi(numberOfDisks-1, bufferPost, toPost, fromPost)

}
