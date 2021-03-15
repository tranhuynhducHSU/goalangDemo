package main

import (
	"fmt"
	"math"
	// "math"
	// "strings"
)

type Vector struct {
	X, Y int
}

func printArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] < 10 {
				fmt.Printf("   %v", arr[i][j])
			} else if arr[i][j] < 100 {
				fmt.Printf("  %v", arr[i][j])
			} else if arr[i][j] < 1000 {
				fmt.Printf(" %v", arr[i][j])
			}
		}
		fmt.Println()
	}
}

func createArr2Way(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	return arr
}

func spiral(arr [][]int) [][]int {
	vectors := createVector()
	vecRight := vectors[1]
	vecTop := vectors[0]

	//trung tâm nè
	mid := 0
	indexX := 0
	indexY := 0
	midCenter := Vector{0, 0}
	if math.Mod(float64(len(arr[0])), 2) == 0 {
		mid = len(arr[0]) / 2
		indexX = mid - 1
		indexY = mid
		midCenter.X = mid - 1
		midCenter.Y = mid
	} else {
		mid = len(arr[0]) / 2
		indexX = mid
		indexY = mid
		midCenter.X = mid
		midCenter.Y = mid
	}

	length := len(arr[0]) * len(arr[0])
	for i := 1; i < length; i++ {
		vecRight = vectors[1]
		if arr[vecRight.Y+indexY][vecRight.X+indexX] == 0 && (vecRight.Y+indexY != midCenter.Y || vecRight.X+indexX != midCenter.X) {
			indexX = vecRight.X + indexX
			indexY = vecRight.Y + indexY
			arr[indexY][indexX] = i
			vectors = rotationVec(vectors)
		} else {
			vecTop = vectors[0]
			indexX = vecTop.X + indexX
			indexY = vecTop.Y + indexY
			arr[indexY][indexX] = i
		}
	}
	return arr
}

func createVector() []Vector {
	vectors := make([]Vector, 4)
	vectors[0] = Vector{-1, 0}
	vectors[1] = Vector{0, -1}
	vectors[2] = Vector{1, 0}
	vectors[3] = Vector{0, 1}
	return vectors
}

func rotationVec(vectors []Vector) []Vector {
	i := vectors[0]
	vectors = vectors[1:]
	vectors = append(vectors, i)
	return vectors
}

func main() {
	n := 0
	fmt.Print("Nhập n là độ rộng ma trận: ")
	fmt.Scan(&n)
	arr := createArr2Way(n)
	arr = spiral(arr)
	printArr(arr)
}
