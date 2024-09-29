import "math"

type AllOne struct {
	index map[string]*Node
	list  DoubleLinkList
}

func Constructor() AllOne {
	index := make(map[string]*Node)
	list := NewDefaultList()
	return AllOne{index: index, list: list}
}

func (this *AllOne) Inc(key string) {
	keyNode, found := this.index[key]
	if !found {
		keyNode = &Node{
			Prev:  nil,
			Next:  nil,
			Value: 0,
			Key:   key,
		}
		this.index[key] = keyNode
		this.list.Insert(keyNode)
	}
	keyNode.Value++
	this.list.SortBackwards(keyNode, this.index)
}

func (this *AllOne) Dec(key string) {
	this.index[key].Value--
	if this.index[key].Value == 0 {
		this.list.Delete(this.index[key], this.index)
	} else {
		this.list.SortForward(this.index[key], this.index)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.list.Head() != nil {
		return this.list.Head().Key
	} else {
		return ""
	}
}

func (this *AllOne) GetMinKey() string {
	if this.list.Tail() != nil {
		return this.list.Tail().Key
	} else {
		return ""
	}
}

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
	Key   string
}

type DoubleLinkList interface {
	Insert(node *Node)
	Delete(node *Node, index map[string]*Node)
	Head() *Node
	SortBackwards(node *Node, index map[string]*Node)
	SortForward(node *Node, index map[string]*Node)
	Tail() *Node
}

type DefaultList struct {
	head *Node
	tail *Node
}

func (d *DefaultList) Head() *Node {
	return d.head.Next
}

func (d *DefaultList) Tail() *Node {
	return d.tail
}

func NewDefaultList() DoubleLinkList {
	dummyNode := &Node{
		Prev:  nil,
		Next:  nil,
		Value: math.MaxInt,
	}
	return &DefaultList{
		head: dummyNode,
		tail: dummyNode,
	}
}

func (d *DefaultList) Insert(node *Node) {
	d.tail.Next = node
	node.Prev = d.tail
	d.tail = node
}

func (d *DefaultList) Delete(node *Node, index map[string]*Node) {
	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if d.tail == node {
		d.tail = node.Prev
	}
	delete(index, node.Key)
}

func (d *DefaultList) swapWithNext(node *Node, index map[string]*Node) {
	index[node.Key] = node.Next
	index[node.Next.Key] = node
	node.Value, node.Next.Value = node.Next.Value, node.Value
	node.Key, node.Next.Key = node.Next.Key, node.Key
}

func (d *DefaultList) swapWithPrevious(node *Node, index map[string]*Node) {
	index[node.Key] = node.Prev
	index[node.Prev.Key] = node
	node.Value, node.Prev.Value = node.Prev.Value, node.Value
	node.Key, node.Prev.Key = node.Prev.Key, node.Key
}

func (d *DefaultList) SortBackwards(node *Node, index map[string]*Node) {
	for node.Value > node.Prev.Value {
		d.swapWithPrevious(node, index)
		node = node.Prev
	}
}

func (d *DefaultList) SortForward(node *Node, index map[string]*Node) {
	for node != nil && node.Next != nil && node.Value < node.Next.Value {
		d.swapWithNext(node, index)
		node = node.Next
	}
}
