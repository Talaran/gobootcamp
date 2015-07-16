package main
import (
	"fmt"
	"github.com/golang/tour/tree"
)

func main() {
	report := make(chan int)
	go walk(tree.New(1), report)
	for value := range report {
		fmt.Println(value)
	}
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(1), tree.New(2)))

}

func walk(tree *tree.Tree, report chan int) {
	recWalk(tree, report)
	close(report)
}

func recWalk(tree *tree.Tree, report chan int) {
	if tree == nil {
		return
	}
	recWalk(tree.Left, report)
	report <- tree.Value
	recWalk(tree.Right, report)
}

func same(firstTree *tree.Tree, secondTree *tree.Tree) bool {
	firstReport := make(chan int)
	secondReport := make(chan int)
	go walk(firstTree, firstReport)
	go walk(secondTree, secondReport)
	for {
		firstValue, firstTest := <-firstReport
		secondValue, secondTest := <-secondReport
		switch {
		case firstTest != secondTest:
			return false
		case ! firstTest:
			return true
		case firstValue != secondValue:
			return false
		}
	}
}