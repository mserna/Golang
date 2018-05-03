package main

import "time"

func quickSort(mArray []int) []int {
	defer timeTrack(time.Now(), "quickSort")
	return mArray
}
