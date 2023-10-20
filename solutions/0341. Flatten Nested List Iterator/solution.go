/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(el T) {
	s.data = append(s.data, el)
}

func (s *Stack[T]) Pop() T {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack[T]) Peek() T {
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

type NestedIterator struct {
	stack *Stack[*NestedInteger]
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := NewStack[*NestedInteger]()
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack.Push(nestedList[i])
	}

	return &NestedIterator{
		stack: stack,
	}
}

func (it *NestedIterator) Next() int {
	return it.stack.Pop().GetInteger()
}

func (it *NestedIterator) HasNext() bool {
	for !it.stack.IsEmpty() {
		top := it.stack.Peek()
		if top.IsInteger() {
			return true
		}

		list := it.stack.Pop().GetList()
		for i := len(list) - 1; i >= 0; i-- {
			it.stack.Push(list[i])
		}
	}

	return false
}
