package main

import (
	"fmt"
	"local/bubblesort"
	"local/selectionsort"
	"local/insertionsort"
	"local/mergesort"
)

func main() {
	unsortedArr :=  []int{ 2,5,1,8,3,9,0}
	sortedArr := bubblesort.SortArr(unsortedArr)
	fmt.Println(sortedArr)
	unsortedArr =  []int{ 2,5,1,8,3,9,0}
	sortedArr = selectionsort.SortArr(unsortedArr)
	fmt.Println(sortedArr)
	unsortedArr =  []int{ 2,5,1,8,3,9,0}
	sortedArr = insertionsort.SortArr(unsortedArr)
	fmt.Println(sortedArr)
	unsortedArr =  []int{ 2,5,1,8,3,9,0}
	sortedArr = mergesort.SortArr(unsortedArr)
	fmt.Println(sortedArr)
}