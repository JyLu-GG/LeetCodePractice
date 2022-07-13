package leetcodepractice

import (
	"fmt"
	"math"
)

func MaxArea1(height []int) int {
	h := 0
	w := 0
	currentarea := 0
	maxarea := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i == j {
				continue
			} else {
				w = j - i
				h = int(math.Min(float64(height[i]), float64(height[j])))
				currentarea = w * h
				if maxarea < currentarea {
					maxarea = currentarea
				}
			}
		}
	}
	return maxarea
}

func MaxArea2(height []int) int {
	h := 0
	w := 0
	currentarea := 0
	maxarea := 0
	index := 0
	lastindex := len(height) - 1
	for index < lastindex {
		w = lastindex - index
		h = int(math.Min(float64(height[index]), float64(height[lastindex])))
		currentarea = w * h
		if maxarea < currentarea {
			maxarea = currentarea
		}
		// Third submit fail reason, use index and lastindex to check not height value
		if height[index] <= height[lastindex] {
			index += 1
		} else {
			lastindex -= 1
		}
	}

	return maxarea
}

func MaxAreaFail(height []int) int {
	h := 0
	w := 0
	currentarea := 0
	maxarea := 0
	index := 0
	lastindex := len(height) - 1

	for i := 0; i < len(height)/2; i++ {
		w = lastindex - index
		h = int(math.Min(float64(height[index]), float64(height[lastindex])))
		currentarea = w * h
		if maxarea < currentarea {
			maxarea = currentarea
		}
		// fmt.Printf("w = %d, h = %d, current = %d, max = %d\n", w, h, currentarea, maxarea)
		// fmt.Printf("indexvalue = %d, lastindexvalue = %d\n", height[index], height[lastindex])

		index += 1
		w = lastindex - index
		h = int(math.Min(float64(height[index]), float64(height[lastindex])))
		currentarea = w * h
		if maxarea < currentarea {
			maxarea = currentarea
		}
		// fmt.Printf("w = %d, h = %d, current = %d, max = %d\n", w, h, currentarea, maxarea)
		// fmt.Printf("indexvalue = %d, lastindexvalue = %d\n", height[index], height[lastindex])

		lastindex -= 1
		w = lastindex - index
		h = int(math.Min(float64(height[index]), float64(height[lastindex])))
		currentarea = w * h
		if maxarea < currentarea {
			maxarea = currentarea
		}
		// fmt.Printf("w = %d, h = %d, current = %d, max = %d\n", w, h, currentarea, maxarea)
		// fmt.Printf("indexvalue = %d, lastindexvalue = %d\n", height[index], height[lastindex])
	}
	fmt.Printf("Fail ans = %d\n", maxarea)
	return maxarea
}
