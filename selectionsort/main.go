package main

// Time complexity O(n * n)
func selectionSort(unorderedList []int) []int {
	listLength := len(unorderedList)

	for i := 0; i < listLength; i++ {
		lowestNumberIndex := i

		for j := i + 1; j < listLength; j++ {
			if unorderedList[j] < unorderedList[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		unorderedList[i], unorderedList[lowestNumberIndex] = unorderedList[lowestNumberIndex], unorderedList[i]
	}

	return unorderedList
}
