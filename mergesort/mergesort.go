package mergesort

func SortArr(inputArr []int) []int {
	startIndex := 0
	endIndex := len(inputArr) - 1
	Mergesort(inputArr, startIndex, endIndex)
	return inputArr
}

func  Mergesort(arr []int , startIndex int, endIndex int)  {
	if(startIndex < endIndex){
		midIndex := (startIndex+endIndex)/2
		Mergesort(arr, startIndex, midIndex)
		Mergesort(arr, midIndex+1, endIndex)
		Merge(arr, startIndex, midIndex, endIndex)
	}
}

func Merge(arr []int, startIndex int, midIndex int, endIndex int) {
	sortedArr := make([]int, endIndex+1)
	sortedIndex := 0
	start := startIndex
	mid := midIndex+1
	for start <= midIndex && mid <= endIndex {
		if(arr[start]<arr[mid]){
			sortedArr[sortedIndex] = arr[start]
			sortedIndex++
			start++
		} else {
			sortedArr[sortedIndex] = arr[mid]
			sortedIndex++
			mid++
		}
	}
	
	for start <= midIndex {
		sortedArr[sortedIndex] = arr[start]
		sortedIndex++
			start++
	}

	for mid<=endIndex {
		sortedArr[sortedIndex] = arr[mid]
		sortedIndex++
			mid++
	}
	sortedIndex=0
	for i := startIndex; i <= endIndex; i++ {
		arr[i] = sortedArr[sortedIndex]
		sortedIndex++
	}
}