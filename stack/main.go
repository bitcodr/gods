package main

import "fmt"

type Node struct {
	value int
}

type Stack struct {
	Node  []*Node
	count int
}

func (s *Stack) push(n *Node) {
	s.Node = append(s.Node, n)
	s.count++
}

func (s *Stack) pop() *Node {
	if s.count == 0 {
		return nil
	}

	s.count--
	return s.Node[s.count]
}

func main() {
	stack := &Stack{}
	stack.push(&Node{value: 4})
	stack.push(&Node{value: 9})
	stack.push(&Node{value: 1})

	fmt.Println(stack.pop(), " - ", stack.pop(), " - ", stack.pop(), " - ")

}
