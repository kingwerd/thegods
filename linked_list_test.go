package thegods

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	got := CreateLinkedList()
	if got == nil {
		t.Error("CreateLinkedList was incorrect, got: ", got)
	}
}

func TestCreateLLNode(t *testing.T) {
	got := CreateLLNode(9)
	if got.Val != 9 {
		t.Errorf("CreateLLNode was incorrect, got %d expected: %d", got.Val, 9)
	}
}

func TestInsertFront(t *testing.T) {
	ll := LinkedList{}
	ll.InsertFront(123)

	// if there is no head insert the node
	if ll.Head.Val != 123 {
		t.Errorf("TestInsertFront was incorrect, got: %d expected: %d", ll.Head.Val, 123)
	}

	ll.InsertFront(321)
	// if there is a head insert the new head
	if ll.Head.Val != 321 {
		t.Errorf("TestInsertFront was incorrect, got: %d expected: %d", ll.Head.Val, 321)
	}
}

func TestInsertBack(t *testing.T) {
	ll := LinkedList{}
	ll.InsertBack(123)

	// if there is no head then the node should be inserted at the head of the list
	if ll.Head.Val != 123 {
		t.Errorf("TestInsertBack was incorrect, got: %d expected: %d", ll.Head.Val, 123)
	}

	ll.InsertBack(9)
	ll.InsertBack(4455)
	// if there is a head move down the list until the final value is found
	runner := ll.Head
	for runner.Next != nil {
		runner = runner.Next
	}

	if runner.Val != 4455 {
		t.Errorf("TestInsertBack was incorrect, got: %d expected: %d", runner.Val, 4455)
	}
}

func TestRemoveFront(t *testing.T) {
	ll := &LinkedList{}
	ll.Head = &LLNode{Val: 123}
	ll.Head.Next = &LLNode{Val: 9999}

	if ok := ll.RemoveFront(); !ok {
		t.Error("RemoveFront failed to remove the front node")
	}
}
