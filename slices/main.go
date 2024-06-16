package main

import (
	"fmt"
	"sort"
)

func demo() {
	var arr = []int{1, 2, 3, 4, 5}

	arr = append(arr, 10, 100)

	arr = append(arr[0:4])

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 300
	highScores[2] = 600
	highScores[3] = 250

	// highScores[4] = 1000

	highScores = append(highScores, 1000, 342)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	var courses = []string{"swift", "rust", "golang", "haskell", "c"}
	fmt.Println(courses)

	var index = 2

	courses = append(courses[0:index], courses[index+1:]...)

	fmt.Println(courses)

	for i, v := range courses {
		fmt.Printf("Index %d has %v\n", i, v)
	}

	var brr = [][]uint{}

	for i := 0; i < 4; i++ {
		brr = append(brr, []uint{1, 2, 4, 5})
	}

	fmt.Println(brr)

}

func main() {
	demo()
}
