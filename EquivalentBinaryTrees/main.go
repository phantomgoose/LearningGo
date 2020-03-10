package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	val1, val2 := make(chan int), make(chan int)
	go Walk(t1, val1)
	go Walk(t2, val2)

	treeVals1, treeVals2 := make([]int, 10), make([]int, 10)
	for i := 0; i < 10; i++ {
		treeVals1[i] = <-val1
		treeVals2[i] = <-val2
	}

	for i, v1 := range treeVals1 {
		if treeVals2[i] != v1 {
			return false
		}
	}
	return true
}

func main() {
	for range [10]int{} {
		fmt.Println(Same(tree.New(1), tree.New(2)))
	}
}
