package selectionsort

func SortArr(inputArr []int) []int {
	startIndex := 0
	for range inputArr {
		minNumIndex := startIndex;
		for i := startIndex+1; i < len(inputArr); i++ {
			if(inputArr[minNumIndex]>inputArr[i]) {
				minNumIndex = i
			}
		}
		inputArr[minNumIndex], inputArr[startIndex] = inputArr[startIndex], inputArr[minNumIndex]
		startIndex+=1
	}
	return inputArr
}