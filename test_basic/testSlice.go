package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = append(scores, 5)
// 	fmt.Println(scores) // prints [5]
// }

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = scores[0:10]
// 	scores[5] = 9033
// 	scores = append(scores, 101)
// 	fmt.Println(scores)
// }

// func main() {
// 	scores := make([]int, 0, 5)
// 	c := cap(scores)
// 	fmt.Println(c)
// 	for i := 0; i < 25; i++ {
// 		scores = append(scores, i)
// 		// if our capacity has changed,
// 		// Go had to grow our array to accommodate the new data
// 		if cap(scores) != c {
// 			c = cap(scores)
// 			fmt.Println(c)
// 		}
// 	}
// }

// func main() {
// 	scores := make([]int, 5)
// 	scores = append(scores, 9332)
// 	fmt.Println(scores)
// }

// func main() {
// 	scores := []int{1, 2, 3, 4, 5}
// 	scores = removeAtIndex(scores, 2)
// 	fmt.Println(scores)
// }

// func removeAtIndex(source []int, index int) []int {
// 	lastIndex := len(source) - 1
// 	//swap the last value and the value we want to remove
// 	source[index], source[lastIndex] = source[lastIndex], source[index]
// 	return source[:lastIndex]
// }


func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:3])
	fmt.Println(worst)
}