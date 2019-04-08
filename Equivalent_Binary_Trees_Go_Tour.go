package main

import (
	"fmt"
	"time"
	"golang.org/x/tour/tree"
)
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t,ch)
	close(ch)
}
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkTree(t *tree.Tree, ch chan int) {
	if t == nil { 
		return
	}
	WalkTree(t.Left, ch)
	ch <-t.Value
	WalkTree(t.Right, ch)
	//fmt.Printf("|%v|,|%v|,|%v|",t.Value,t.Left,t.Right)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i:=0;i<10;i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	//fmt.Printf("%v",tree.New(1))
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Print("->",i," ")
	}
	fmt.Printf("\n%v",Same(tree.New(1), tree.New(1)))
	fmt.Printf("\n%v",Same(tree.New(1), tree.New(2)))
	time.Sleep(time.Millisecond)
}