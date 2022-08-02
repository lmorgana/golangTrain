package main

import (
	"fmt"
	"math/rand"
)

func randomInt(max, min int) int {
	return rand.Intn(max-min) + min
}

func makeArr3d(x, y, z int) [][][]int {
	arr := make([][][]int, x)
	for i, _ := range arr {
		arr[i] = make([][]int, y)
		for j, _ := range arr[i] {
			arr[i][j] = make([]int, z)
			for k, _ := range arr[i][j] {
				arr[i][j][k] = randomInt(10, 0)
			}
		}
	}
	return arr
}

func printArr(arr [][][]int) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func sumArray3d(arrA, arrB [][][]int) [][][]int {
	arr := makeArr3d(len(arrA), len(arrA[0]), len(arrA[0][0]))
	for i, _ := range arr {
		for j, _ := range arr[i] {
			for k, _ := range arr[i][j] {
				arr[i][j][k] = arrA[i][j][k] + arrB[i][j][k]
			}
		}
	}
	return arr
}

func subArray3d(arrA, arrB [][][]int) [][][]int {
	arr := makeArr3d(len(arrA), len(arrA[0]), len(arrA[0][0]))
	for i, _ := range arr {
		for j, _ := range arr[i] {
			for k, _ := range arr[i][j] {
				arr[i][j][k] = arrA[i][j][k] - arrB[i][j][k]
			}
		}
	}
	return arr
}

func main() {
	arrA := makeArr3d(2, 2, 2)
	arrB := makeArr3d(2, 2, 2)

	fmt.Println("ArrayA:")
	printArr(arrA)
	fmt.Println("ArrayB:")
	printArr(arrB)

	arrSum := subArray3d(arrA, arrB)
	fmt.Println("Array:")
	printArr(arrSum)
}
