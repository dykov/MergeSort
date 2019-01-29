package main

import "fmt"

func main() {

	var array = []int{10,2,1,100,23,12}
	fmt.Println( MergeSort(array) )
	
}

func MergeSort(array []int) []int {

	if len(array) < 2 {
		return array
	}
	middle := (len(array)) / 2
	return Merge(MergeSort(array[:middle]), MergeSort(array[middle:]))
}

func Merge(left, right []int) []int {

	size := len(left)+len(right)
	i := 0
	j := 0
	result := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			result[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}