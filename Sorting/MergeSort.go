package main

import "time"

func merge(arr []int, left int, mid int, right int) {

	i := 0
	j := 0
	k := 1

	n1 := mid - left + 1 //Pointer 1
	n2 := right - mid    //Pointer 2

	//Temp arrays
	leftArray := make([]int, n1)
	rightArray := make([]int, n2)

	//Merge temp arrays back into arr[1..n]
	for ; i < n1; i++ {
		leftArray[i] = arr[left+i]
	}
	for ; j < n2; j++ {
		rightArray[j] = arr[mid+left+j]
	}
	for i < n1 && j < n2 {
		if leftArray[i] <= rightArray[j] {
			arr[k] = leftArray[i]
			i++
		} else {
			arr[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArray[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArray[j]
		j++
		k++
	}
}

func mergeSort(mArray []int, left int, right int) []int {
	defer timeTrack(time.Now(), "mergeSort")

	if left < right {
		mid := left + (right-left)/2
		mergeSort(mArray, left, mid)
		mergeSort(mArray, mid+1, right)

		merge(mArray, left, mid, right)
	}
	return mArray
}
