package towerofhanoi

import (
	"fmt"
	"testing"
)

func TestHanoiTower(t *testing.T) {

	fmt.Println("Move 3 disks from post 1 to post 3")
	towerOfHanoi(3, 1, 3, 2)
	fmt.Println("\nMove 10 disks from post 1 to post 3")
	towerOfHanoi(10, 1, 3, 2)

}
