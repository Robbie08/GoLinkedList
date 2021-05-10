package main

// create our node
type node struct {
	data int   // attribute date that will be of type int
	next *node // attribute next which will be a pointer to the next node
}

type linkedList struct {
	oHead *node
	oTail *node
	oSize int
}
