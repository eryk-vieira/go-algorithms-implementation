package main

// Time complexity O(log₂n)
// This means, in the worst case the target requires log₂(n). where n is the size of the list.
func binarySearch(orderedList []int, search int) *int {
	down := 0
	up := len(orderedList) - 1

	for down <= up {
		middle := (down + up) / 2
		guess := orderedList[middle]

		if guess == search {
			return &middle
		}

		if guess > search {
			up = middle - 1
		} else {
			down = middle + 1
		}
	}

	return nil
}

// Time complexity O(n)
// This means, in the worst case will need to scan all the list in order to find or not the target.
func linearSearch(orderedList []int, search int) *int {
	for i, item := range orderedList {
		if item == search {
			return &i
		}
	}

	return nil
}
