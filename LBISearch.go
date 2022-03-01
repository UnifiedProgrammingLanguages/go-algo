package main

import "fmt"

/* Linear Search
------------------------------------
This technique pass over the list of elements, by using the index to move from the beginning of the list to the end. Each element is examined and if it does not match the search item, the next item is examined.
By hopping from one item to its next, the list is passed over sequentially.
------------------------------------
*/
func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

/* Binary Search
------------------------------------
A binary search is a search strategy used to find elements within a list by consistently reducing the amount of data to be searched and thereby increasing the rate at which the search term is found.
To use a binary search algorithm, the list to be operated on must have already been sorted.
------------------------------------
*/
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

/* Interpolation Search
------------------------------------
The Interpolation Search is an improvement over Binary Search for instances, where the values in a sorted array are uniformly distributed. Binary Search always goes to middle element to check.
On the other hand interpolation search may go to different locations according the value of key being searched. Here is the source code of the Go program to search element in an integer array using Interpolation search algorithm.
The output shows the position of element in array.
------------------------------------
*/
func interpolationSearch(array []int, key int) int {

	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if array[guess] == key {
			// scan backwards for start of value range
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	//Linear Search
	fmt.Println(linearsearch(items, 63))

	//Binary Search
	fmt.Println(binarySearch(63, items))

	//Interpolation Search
	fmt.Println(interpolationSearch(items, 63))
}
