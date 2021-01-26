package thegods

// LLNode holds an integer Val and points to the
// Next LLNode in the LinkedList.
type LLNode struct {
	Next *LLNode
	Val  int
}

// LinkedList is a data structure that has a head pointing
// to a LLNode or if its empty then nil.
type LinkedList struct {
	Head *LLNode
	Name string
}

// CreateLinkedList initializes an empty LinkedList and then
// returns the LinkedList.
func CreateLinkedList() *LinkedList {
	return &LinkedList{}
}

// CreateLLNode initializes a LLNode with an integer value
// and returns the LLNode.
func CreateLLNode(value int) *LLNode {
	return &LLNode{Val: value}
}

// InsertFront creates a new LLNode and adds the newly
// created LLNode to the front of the LinkedList.
func (ll *LinkedList) InsertFront(value int) {
	if ll.Head == nil {
		ll.Head = CreateLLNode(value)
		return
	}
	newNode := CreateLLNode(value)
	newNode.Next = ll.Head
	ll.Head = newNode
}

// InsertBack creates a new LLNode and adds the newly
// created LLNode to the end of the LinkedList.
func (ll *LinkedList) InsertBack(value int) {
	if ll.Head == nil {
		ll.Head = CreateLLNode(value)
		return
	}
	runner := ll.Head
	for runner.Next != nil {
		runner = runner.Next
	}
	runner.Next = CreateLLNode(value)
}

// RemoveFront remove the LLNode at the Head of the LinkedList
func (ll *LinkedList) RemoveFront() bool {
	ll.Head = ll.Head.Next
	return true
}

// RemoveAtPosition removes the LLNode that is located at
// a position relative to the length of the LinkedList
func (ll *LinkedList) RemoveAtPosition(position int) {
	if ll.Head != nil {
		currentNode := ll.Head
		if position == 0 {
			ll.Head = currentNode.Next
			return
		}
		counter := 0
		previousNode := ll.Head
		for currentNode != nil && counter != position {
			previousNode = currentNode
			currentNode = currentNode.Next
			counter++
		}

		if currentNode == nil {
			return
		}

		previousNode.Next = currentNode.Next
		currentNode = nil
	}
}
