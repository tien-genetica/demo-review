package main

// BinarySearch performs binary search on a sorted array.
// It returns the index of the target if found, otherwise returns -1.
// Time complexity: O(log n)
// Space complexity: O(1)
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2 // Avoid potential overflow

		// Check if target is present at mid
		if arr[mid] == target {
			return mid
		}

		// If target is greater, ignore left half
		if arr[mid] < target {
			left = mid + 1
		} else {
			// If target is smaller, ignore right half
			right = mid - 1
		}
	}

	// Target not found
	return -1
}

// BinarySearchRecursive is a recursive implementation of binary search.
// It returns the index of the target if found, otherwise returns -1.
// Time complexity: O(log n)
// Space complexity: O(log n) due to recursion stack
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}

	return BinarySearchRecursive(arr, target, mid+1, right)
}
