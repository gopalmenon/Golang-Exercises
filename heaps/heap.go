package heaps

import "fmt"

func parentIndex(nodeIndex int) int {
	return (nodeIndex - 1) / 2
}

func leftChildIndex(nodeIndex int) int {
	return 2*nodeIndex + 1
}

func rightChildIndex(nodeIndex int) int {
	return leftChildIndex(nodeIndex) + 1
}

func maxHeapify(array []int, rootIndex, heapSize int) {

	l, r := leftChildIndex(rootIndex), rightChildIndex(rootIndex)

	largest := 0
	if l < heapSize && array[l] > array[rootIndex] {
		largest = l
	} else {
		largest = rootIndex
	}

	if r < heapSize && array[r] > array[largest] {
		largest = r
	}

	if largest != rootIndex {
		array[rootIndex], array[largest] = array[largest], array[rootIndex]
		maxHeapify(array, largest, heapSize)
	}

}

func BuildMaxHeap(array []int) {

	for index := len(array)/2 - 1; index >= 0; index-- {
		maxHeapify(array, index, len(array))
	}

}

func HeapSort(array []int) {

	BuildMaxHeap(array)
	heapSize := len(array)
	for index := len(array) - 1; index >= 1; index-- {
		array[0], array[index] = array[index], array[0]
		heapSize -= 1
		maxHeapify(array, 0, heapSize)
	}

}

func String(array []int) {

	for index, element := range array {

		if (index+1)&index == 0 {
			fmt.Print("\n")
		}

		fmt.Printf("\t%d", element)

	}
	fmt.Print("\n")
}
