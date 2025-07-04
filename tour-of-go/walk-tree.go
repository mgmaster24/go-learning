package gol_tourofgo

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1c := make(chan int)
	t2c := make(chan int)
	Walk(t1, t1c)
	Walk(t2, t2c)

	return <-t1c == <-t2c
}
