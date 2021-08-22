package main

import "fmt"

type Node struct {
	value int
}

type Queue struct {
	nodes []*Node
}

func (q *Queue) push(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) pop() *Node {
	if len(q.nodes) == 0 {
		return nil
	}

	qu := q.nodes[0]

	q.nodes = q.nodes[1:]

	return qu
}

func main() {
 	qq := &Queue{}
 	qq.push(&Node{value: 9})
	qq.push(&Node{value: 8})
	qq.push(&Node{value: 3})

	fmt.Println(qq.pop(), " - ", qq.pop(), " - ", qq.pop(), " - ")


}
