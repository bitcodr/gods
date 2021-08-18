package main

import "fmt"

type node struct {
	data int
	next *node
	//prev *node  for double direction linked list
}

func newNode(x int) *node {
	n := &node{
		data: x,
	}

	return n
}

func numberOfNodes(head *node) int {
	x := 1
	current := head

	for current.next != nil {
		current = current.next
		x++
	}

	return x
}

func main() {
	var linkedList []*node

	a, c, f, g := 1, 3, 5, 7

	head := newNode(c)
	nexta := newNode(a)
	nextb := newNode(f)
	nextc := newNode(g)

	head.next = nexta
	nexta.next = nextb
	nextb.next = nextc

	linkedList = append(linkedList, head, nexta, nextb, nextc)

	for _, v := range linkedList {
		if v.next != nil {
			fmt.Println(v.data, " - ", v.next.data)
		} else {
			fmt.Println(v.data, " - ", nil)
		}
	}

	fmt.Println(numberOfNodes(head))
}
