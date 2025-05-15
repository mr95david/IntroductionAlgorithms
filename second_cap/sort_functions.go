package main

import (
	"fmt"

	"time"
)

// This function consider the algorithm to increasing and decreasing sort
func InsertionSort(input_arr []int, increasing bool) {
	start_time := time.Now()
	fmt.Println("This is a function to sort the list: ", input_arr)
	// Ciclo para recorrer cada uno de los valores dentro del array (Considerando que en go, el valor inicial es 0)
	for j := 1; j < len(input_arr); j += 1 {
		// Primero se toma el valor llave que se busca evaluar "key"
		var key int = input_arr[j]
		var i int = j - 1

		// Se un ciclo para evaluar del primer valor a la derecha, hasta el primer valor del array para verificar si el valor llave es menor a al valor i que se esta evaluando
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

		// Se asigna el valor a la ultima posicion evaluada de la llave
		input_arr[i+1] = key
	}
	total_duration := time.Since(start_time)
	fmt.Println("The new array, currently sorted is: ", input_arr)
	fmt.Println("Time of Duration: ", total_duration)
}
