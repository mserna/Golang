package main

import (
	"fmt"
	"time"
)

func printArray(pArray []int) {
	for index := 0; index < len(pArray); index++ {
		fmt.Printf("%d ", pArray[index])
	}
	fmt.Printf("\n")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s: ", name, elapsed)
}

func main() {

	fmt.Printf("Sorting using different methods.")
	fmt.Println()

	array := []int{333, 3, 22, 12., 72, 811}
	fmt.Printf("Original array: ")
	printArray(array)
	fmt.Println()

	//Bubble Sort
	bSortArray := bubbleSort(array)
	fmt.Printf("Bubble sorted array: ")
	printArray(bSortArray)
	fmt.Println()

	//Merge Sort
	mSortArray := mergeSort(array, 0, len(array)-1)
	fmt.Printf("Merge sorted array: ")
	printArray(mSortArray)
	fmt.Println()

	//Quick Sort
	qSortArray := quickSort(array)
	fmt.Printf("Quick sorted array: ")
	printArray(qSortArray)
	fmt.Println()
}
