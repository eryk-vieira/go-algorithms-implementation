package main

// Time complexity O(log₂n)
// This means, in the worst case the target requires log₂(n). where n is the size of the list.
func binarySearch(orderedList []int, search int) int {
	low := 0
	high := len(orderedList) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := orderedList[mid]

		if guess == search {
			return mid
		}

		if guess > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// Time complexity O(n)
// This means, in the worst case will need to scan all the list in order to find or not the target.
// Returns -1 if the search was not found
func linearSearch(orderedList []int, search int) int {
	for i := 0; i < len(orderedList); i++ {
		if search == orderedList[i] {
			return i
		}
	}

	return -1
}
