package main

import "fmt"

func swap(array []int, x int, y int) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}

func bubbleSort(mArray []int) []int {
	length := len(mArray)
	for i := length - 2; i > 1; i-- {
		for j := 0; j < i; j++ {
			if mArray[j] > mArray[j + 1] {
				swap(mArray, j, j + 1)
			}
		}
	}

	return mArray
}

func printArray(pArray []int) {
	for index := 0; index < len(pArray); index++ {
		fmt.Printf("%d ", pArray[index])
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Printf("Bubble Sort!\n")

	array := []int{333, 3, 22, 12., 72, 811}
	fmt.Printf("Original array: ")
	printArray(array)

	sortedArray := bubbleSort(array)
	fmt.Printf("Sorted array: ")
	printArray(sortedArray)
}