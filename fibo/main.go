package main

import (
	// "strings"
	// "golang.org/x/tour/wc"
	"fmt"
)

// func WordCount(s string) (dict map[string]int) {
// 	words := strings.Split(s, " ")
// 	dict = make(map[string]int)
// 	for _, word := range words {
// 		wordC := dict[word]
// 		dict[word] = wordC + 1
// 	}
// 	return dict
// }
func fibonacci() func() int {
	x, y := -1, 0
	return func() int {
		t := x
		x = y
		y = t + y
		return -x
	}

}

func main() {
	//wc.Test(WordCount)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
