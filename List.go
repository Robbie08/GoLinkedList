package main

type List interface {
	add(int) bool // adds the value to the start of the list

	addLast(int) bool // adds the value to the end of the list

	remove() *node // removes the value at the front of the list

	delete(int) bool // removes all instances of the value passed in

	getHeadVal() int // will return the value at the head of list

	getTailVal() int // will return the value at the tail of list

	getSize() int // will retun the number of items in the list

	isEmpty() bool // will return true if the list is empty

	contains(int) bool // will return true if the value is in the list

	printList() // prints the entire list
}
