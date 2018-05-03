package main

import "time"

func swap(array []int, x int, y int) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}

func bubbleSort(mArray []int) []int {
	defer timeTrack(time.Now(), "bubbleSort")
	length := len(mArray)
	for i := length - 2; i > 1; i-- {
		for j := 0; j < i; j++ {
			if mArray[j] > mArray[j+1] {
				swap(mArray, j, j+1)
			}
		}
	}

	return mArray
}
