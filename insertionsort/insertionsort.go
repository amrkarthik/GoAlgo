package insertionsort

func SortArr(inputArr []int) []int {
	for i := 1; i < len(inputArr); i++ {
		value := inputArr[i]
		minNumIndex := i -1
		for minNumIndex >= 0 && value<inputArr[minNumIndex] {
			inputArr[minNumIndex+1]=inputArr[minNumIndex]
			minNumIndex--
		}
		inputArr[minNumIndex+1] = value
	}
	return inputArr
}