/*
* This program will implement a basic linked list using Golang
* Author: Robert Ortiz
* Date  : 5/9/2021
 */

package main

import (
	"fmt"
)

func main() {
	// example of how to interact with the List interface
	var list List = &linkedList{} // implementing the interface with a LinkedList
	list.add(12)
	list.add(29)
	list.add(10)
	list.addLast(55)
	list.add(2000)
	list.add(1986)
	list.addLast(215)
	list.add(10)
	list.addLast(99)
	list.delete(99)
	list.delete(10) // removes all the instances of the value 10
	list.update(1, -8)

	fmt.Println()
	fmt.Println("////////////////////////////////")
	fmt.Println("////   LinkedList - Tests   ////")
	fmt.Println("////////////////////////////////")

	// print out the list information
	fmt.Println("Head: ", list.getHeadVal())
	fmt.Println("Tail: ", list.getTailVal())
	fmt.Println("List Size: ", list.getSize())
	fmt.Println("contains(29): ", list.contains(29))
	fmt.Println("contains(99): ", list.contains(99))
	fmt.Println("contains(12): ", list.contains(12))
	list.printList()
	fmt.Println("\nisEmpty(): ", list.isEmpty())
	fmt.Println()
}
