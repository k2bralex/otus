package hw_4

/*
	Realize Doubly Linked List structure
*/

type Node struct {
	pNext, pPrev *Node
	data         int
}

func (n *Node) GetNext() *Node {
	return n.pNext
}

func (n *Node) GetPrev() *Node {
	return n.pPrev
}

func (n *Node) GetData() int {
	return n.data
}

type DoublyLinkedList struct {
	head, tail *Node
}

func (dl *DoublyLinkedList) GetHead() *Node {
	return dl.head
}

func (dl *DoublyLinkedList) GetTail() *Node {
	return dl.tail
}

func (dl *DoublyLinkedList) PushFront(data int) *Node {
	nodePtr := &Node{data: data}
	nodePtr.pNext = dl.head
	if dl.head != nil {
		dl.head.pPrev = nodePtr
	}
	if dl.tail == nil {
		dl.tail = nodePtr
	}
	dl.head = nodePtr
	return nodePtr
}

func (dl *DoublyLinkedList) PushBack(data int) *Node {
	nodePtr := &Node{data: data}
	nodePtr.pPrev = dl.tail
	if dl.tail != nil {
		dl.tail.pNext = nodePtr
	}
	if dl.head == nil {
		dl.head = nodePtr
	}
	dl.tail = nodePtr
	return nodePtr
}

func (dl *DoublyLinkedList) PopFront() {
	if dl.head == nil {
		return
	}
	nodePtr := dl.head.pNext
	if nodePtr != nil {
		nodePtr.pPrev = nil
	} else {
		dl.tail = nil
	}
	dl.head = nodePtr
}

func (dl *DoublyLinkedList) PopBack() {
	if dl.tail == nil {
		return
	}
	nodePtr := dl.head.pPrev
	if nodePtr != nil {
		nodePtr.pNext = nil
	} else {
		dl.head = nil
	}
	dl.tail = nodePtr
}

func (dl *DoublyLinkedList) GetAT(index int) *Node {
	nodePtr := dl.head
	n := 0
	for n != index {
		if nodePtr == nil {
			return nodePtr
		}
		nodePtr = nodePtr.pNext
		n++
	}

	return nodePtr
}

func (dl *DoublyLinkedList) Insert(index, data int) *Node {
	nodePtrRight := dl.GetAT(index)
	if nodePtrRight == nil {
		return dl.PushBack(data)
	}
	nodePtrLeft := nodePtrRight.pPrev
	if nodePtrLeft == nil {
		return dl.PushFront(data)
	}
	nodePtr := &Node{data: data}
	nodePtr.pPrev = nodePtrLeft
	nodePtr.pNext = nodePtrRight
	nodePtrLeft.pNext = nodePtr
	nodePtrRight.pPrev = nodePtr

	return nodePtr
}

func (dl *DoublyLinkedList) delete(index int) {
	nodePtr := dl.GetAT(index)
	if nodePtr == nil {
		return
	}
	if nodePtr.pPrev == nil {
		dl.PopFront()
		return
	}
	if nodePtr.pNext == nil {
		dl.PopBack()
		return
	}
	nodeLeft := nodePtr.pPrev
	nodeRight := nodePtr.pNext
	nodeLeft.pNext = nodeRight
	nodeRight.pPrev = nodeLeft
}
