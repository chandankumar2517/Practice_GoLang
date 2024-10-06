package mergelinkedlist

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value int) {

	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head

		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
}

// MergeTwoLists merges two sorted linked lists and returns the head of the new list
func MergeTwoLists(l1, l2 *Node) *Node {

	dummy := &Node{}

	current := dummy

	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			current.next = l1
			l1 = l1.next
		} else {
			current.next = l2
			l2 = l2.next
		}
	}

	current = current.next

	if l1 != nil {
		current.next = l1
	}
}
