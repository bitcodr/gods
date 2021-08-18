package main

import "fmt"

const arraySize = 10

type HashTable struct {
	arr [arraySize]*Bucket
}

type Bucket struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

//hashTable
func (h *HashTable) insert(key string) {
	index := hash(key)
	h.arr[index].insert(key)
}

func (h *HashTable) delete(key string) {
	index := hash(key)
	h.arr[index].delete(key)
}

func (h *HashTable) search(key string) bool {
	index := hash(key)
	return h.arr[index].search(key)
}

//bucket
func (b *Bucket) insert(key string) {
	if !b.search(key) {
		head := b.head
		b.head = &Node{
			data: key,
			next: head,
		}

		return
	}

	fmt.Println("the key already exist")
}

func (b *Bucket) delete(key string) {
	if b.head.data == key {
		b.head = b.head.next
		return
	}

	prev := b.head

	for prev.next != nil {
		if prev.next.data == key {
			prev.next = prev.next.next
			return
		}

		prev = prev.next
	}
}

func (b *Bucket) search(key string) bool {
	current := b.head

	for current != nil {
		if current.data == key {
			return true
		}

		current = current.next
	}

	return false
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}

	return sum % arraySize
}

func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.arr {
		hashTable.arr[i] = &Bucket{}
	}

	return hashTable
}

func main() {
	hashTable := Init()

	keys := []string{"hello", "ok", "thatsGood", "ko"}

	for _, v := range keys {
		hashTable.insert(v)
	}

	response := hashTable.search("ok")
	fmt.Println(response)

	response = hashTable.search("ko")
	fmt.Println(response)

	hashTable.delete("ok")
	response = hashTable.search("ok")
	fmt.Println(response)

	response = hashTable.search("hello")
	fmt.Println(response)

	response = hashTable.search("dcsdc")
	fmt.Println(response)

	hashTable.insert("ko")

}
