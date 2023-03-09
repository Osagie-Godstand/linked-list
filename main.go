package main

import "fmt"

//	All data structures
//
// Heaps(Arrays, list or slice) has a maxheap or minheap and forms a tree structure to demonstrate the relationship
// between the keys. The algo for heap is heapsort which is a sorting algorithm.

// Linked lists stores data in a sequence of nodes and they are linked to eachother using an address that indicates where the other node is.
// Linked list allows you to add and delete something in the beginning element of a list easily.
// A doubly linked list contains the address of the next node and the previus ndoe.
type node struct { // linked list of nodes.
	data int
	next *node // * is a pointer to the next node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) { // prepend means to add something to the beginning of something else.
	second := l.head
	l.head = n
	l.head.next = second // this is a doubly linked list.
	l.length++
}

func (l linkedlist) printlistData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 29}
	node2 := &node{data: 27}
	node3 := &node{data: 16} // prepend makes the address of node3.
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.printlistData()
}
