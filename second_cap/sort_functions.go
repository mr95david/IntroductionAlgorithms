package main

import (
	"fmt"
	"time"
)

// START OF SECTION 2.1
// This function consider the algorithm to increasing and decreasing sort
func InsertionSort(input_arr []int, increasing bool) {
	start_time := time.Now()
	fmt.Println("This is a function to sort the list: ", input_arr)

	for j := 1; j < len(input_arr); j += 1 {

		var key int = input_arr[j]
		var i int = j - 1

		if increasing {
			for i >= 0 && input_arr[i] > key {
				input_arr[i+1] = input_arr[i]
				i = i - 1
			}
		} else {
			for i >= 0 && input_arr[i] < key {
				input_arr[i+1] = input_arr[i]
				i = i - 1
			}
		}

		input_arr[i+1] = key
	}

	total_duration := time.Since(start_time)
	fmt.Println("The new array, currently sorted (By Insertion Sort) is: ", input_arr)
	fmt.Println("Time of Duration: ", total_duration)
}

func LinearSearch(input_arr []int, key int) LinearSearchResult {
	var response LinearSearchResult
	response.exist_validation = false

	for j := 0; j < len(input_arr); j++ {
		if key == input_arr[j] {
			response.value_position = j
			response.exist_validation = true
		}
	}

	if !response.exist_validation {
		fmt.Printf("The value %d doesn't exist in array", key)
	} else {
		fmt.Printf("The value %d exist in array in position %d", key, response.value_position)
	}

	return response
}

// FINAL OF SECTION 2.1

// START OF SECTION 2.3 MERGESORT
func Merge(input_arr []int, p int, q int, r int) {
	// N_n values defined
	var n_1 int = q - p
	var n_2 int = r - q

	// inicialices array copies
	L_arr := make([]int, n_1+1)
	R_arr := make([]int, n_2+1)

	for i := range n_1 {
		L_arr[i] = input_arr[p+i]
	}
	for j := range n_2 {
		R_arr[j] = input_arr[q+j]
	}

	L_arr[len(L_arr)-1] = int_max
	R_arr[len(R_arr)-1] = int_max

	i := 0
	j := 0

	for k := p; k < r; k++ {
		if L_arr[i] <= R_arr[j] {
			input_arr[k] = L_arr[i]
			i++
		} else {
			input_arr[k] = R_arr[j]
			j++
		}
	}
}

func MergeSort(input_arr []int, p int, r int) {
	// fmt.Println(input_arr)

	if p < r-1 {
		var q int

		if r&1 != 0 {

			q = (p+r)/2 + 1
		} else {

			q = (p + r) / 2
		}

		MergeSort(input_arr, p, q)
		MergeSort(input_arr, q, r)
		Merge(input_arr, p, q, r)
	}
}

// END SECTION 2.3
