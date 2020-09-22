package bubblesort

func SortArr(inputArr []int) []int {
	lastUnsortedIndex := len(inputArr) - 1
	isSwapped := false
	for ok := true; ok; ok = isSwapped {
		isSwapped = false
		for i := 0; i < lastUnsortedIndex ; i++ {
			if(inputArr[i] > inputArr[i+1]) {
				inputArr[i+1], inputArr[i] = inputArr[i], inputArr[i+1]
				isSwapped = true
			}
		}
	}
	return inputArr
}