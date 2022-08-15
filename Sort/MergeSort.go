package main

import "fmt"

func Merge(auxArray []int, array []int, low int, mid int, high int) {
	leftIndex := low
	rightIndex := mid + 1
	auxIndex := low

	for leftIndex <= mid && rightIndex <= high {
		if array[leftIndex] < array[rightIndex] {
			auxArray[auxIndex] = array[leftIndex]
			auxIndex++
			leftIndex++
		} else {
			auxArray[auxIndex] = array[rightIndex]
			auxIndex++
			rightIndex++
		}
	}

	if leftIndex <= mid {
		for i := leftIndex; i <= mid; i++ {
			auxArray[auxIndex] = array[i]
			auxIndex++
		}
	} else {
		for i := rightIndex; i <= high; i++ {
			auxArray[auxIndex] = array[i]
			auxIndex++
		}
	}

	for i := low; i <= high; i++ {
		array[i] = auxArray[i]
	}
}

func MergeSort(auxArray []int, array []int, low int, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(auxArray, array, low, mid)
		MergeSort(auxArray, array, mid+1, high)
		Merge(auxArray, array, low, mid, high)
	}
}

func main() {
	array := []int{7, 3, 5, 2, 4, 1, 8, 6, 0, 10, 9}
	fmt.Println("Unsorted array ->", array)

	auxArray := make([]int, len(array))
	MergeSort(auxArray, array, 0, len(array)-1)
	fmt.Println("Sorted array using merge sort ->", array)
}
