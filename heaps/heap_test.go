package heaps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {

	array := []int{9, 7, 8, 5, 6, 4, 3, 2, 0, 1}
	maxHeapify(array, 0, len(array))
	fmt.Printf("Max Heap:\n")
	String(array)
	sortedArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	HeapSort(array)
	if !reflect.DeepEqual(array, sortedArray) {
		t.Errorf("Expected %v but got %v\n", sortedArray, array)
	}
	fmt.Printf("Sorted array:\n")
	String(array)

}
