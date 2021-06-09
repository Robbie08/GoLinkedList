package main

import "fmt"

/*
* This function will allow us to add nodes to the front of the linked list.
*
* in Golang, the function definiton is
* func ({reciever}) {nameOfFunction} ({params}) {}
 */
func (list *linkedList) add(val int) bool {
	n := &node{data: val}
	if list.oHead == nil { // if the list is empty
		list.oHead = n
		list.oTail = n
	} else {
		temp := list.oHead
		list.oHead = n
		list.oHead.next = temp
	}
	list.oSize++ // increase our length
	return true
}

/*
* This function adds a node to the tail of the list
 */
func (list *linkedList) addLast(val int) bool {
	n := &node{data: val}
	if list.oHead == nil {
		list.oHead = n
		list.oTail = n
	} else {
		list.oTail.next = n
		list.oTail = n
	}
	list.oSize++
	return true
}

/*
* This function will remove the node at the front of the list
 */
func (list *linkedList) remove() *node {
	var temp *node
	if list.oHead == nil {
		return nil
	} else {
		temp = list.oHead
		list.oHead = temp.next
	}
	list.oSize--
	return temp
}

func (list *linkedList) delete(val int) bool {
	if list.oHead == nil {
		return false
	}

	var isDeleted bool // this will act as the flag for the deletion

	// if the item is at the start of the list
	if list.oHead.data == val {
		temp := list.oHead
		list.oHead = temp.next
		list.oSize--
		isDeleted = true // set our flag that some item was deleted
	}

	curr := list.oHead.next
	prev := list.oHead
	count := 0
	for curr != list.oTail && count < list.oSize-1 {
		if curr.data == val {
			prev.next = curr.next
			if curr.next != nil {
				curr = curr.next
			}
			list.oSize--
			isDeleted = true
		} else {
			curr = curr.next
			prev = prev.next
		}
		count++
	}

	if list.oTail.data == val {
		list.oTail = prev
		list.oTail.next = nil
		list.oSize--
		isDeleted = true
	}

	return isDeleted // will return if we deleted an item or not
}

func (list *linkedList) update(index, val int) bool {
	if list.isEmpty() {
		return false
	}

	count := 0
	curr := list.oHead
	isUpdated := false
	for count < index && curr != nil {
		curr = curr.next
		count++
	}

	if count == index {
		curr.data = val
		isUpdated = true
	}

	return isUpdated
}

// returns the Head node's data
func (list *linkedList) getHeadVal() int {
	return list.oHead.data
}

// returns the Tail node's data
func (list *linkedList) getTailVal() int {
	return list.oTail.data
}

// returns the number of items in the list
func (list *linkedList) getSize() int {
	return list.oSize
}

func (list *linkedList) isEmpty() bool {
	return list.oHead == nil
}

func (list *linkedList) contains(val int) bool {
	curr := list.oHead

	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

/*
* This function will clear our list
 */

func (list *linkedList) clear() {
	list.oHead = nil
	list.oTail = nil
	list.oSize = 0
}

/*
* This function will print all the elements of the list
 */
func (list *linkedList) printList() {
	curr := list.oHead
	fmt.Print("List: ")
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
}
