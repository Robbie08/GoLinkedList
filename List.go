package main

type List interface {
	add(val int) bool // adds the value to the start of the list

	addLast(val int) bool // adds the value to the end of the list

	remove() *node // removes the value at the front of the list

	delete(val int) bool // removes all instances of the value passed in

	update(index, val int) bool // will update the value at the desired index return false if no index

	getHeadVal() int // will return the value at the head of list

	getTailVal() int // will return the value at the tail of list

	getSize() int // will retun the number of items in the list

	isEmpty() bool // will return true if the list is empty

	contains(val int) bool // will return true if the value is in the list

	clear()

	printList() // prints the entire list
}
