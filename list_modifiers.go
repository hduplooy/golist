package list

// Remove removes element elm from list.
// If the element is nil then nil is returned
// The value of the element is returned.
func (lst *List) Remove(elm *Element) interface{} {
	if elm == nil {
		return nil
	}
	if elm.list == lst {
		lst.remove(elm)
	}
	return elm.Value
}

// RemoveAt removed the element at position pos in the list and return its value
// If position is negative the element from the back of the list will be removed.
// If pos is larger than the length of the list nil is returned.
func (lst *List) RemoveAt(pos int) interface{} {
	elm := lst.Nth(pos)
	if elm == nil {
		return nil
	}
	return lst.Remove(elm)
}

// PushFront inserts a new element with value at the front of list and returns the list.
func (lst *List) PushFront(value interface{}) *List {
	lst.lazyInit()
	lst.insertValue(value, &lst.root)
	return lst
}

// PushBack inserts a new element with value at the back of list lst and returns the list.
func (lst *List) PushBack(value interface{}) *List {
	lst.lazyInit()
	lst.insertValue(value, lst.root.prev)
	return lst
}

// InsertBefore inserts a new element with value immediately before mark and returns the list.
// If mark is not an element of lst or nil, the list is not modified.
func (lst *List) InsertBefore(value interface{}, mark *Element) *List {
	if mark == nil || mark.list != lst {
		return nil
	}
	// see comment in List.Remove about initialization of l
	lst.insertValue(value, mark.prev)
	return lst
}

// InsertAfter inserts a new element with value immediately after mark and returns the list.
// If mark is not an element of lst or nil, the list is not modified.
func (lst *List) InsertAfter(value interface{}, mark *Element) *List {
	if mark == nil || mark.list != lst {
		return nil
	}
	// see comment in List.Remove about initialization of l
	lst.insertValue(value, mark)
	return lst
}

// InsertListAfter inserts a new list into the current list lst after mark and return the current list
// The other list is not modified. If either the other list or mark is nil or mark is not part of lst then lst is returned.
func (lst *List) InsertListAfter(other *List, mark *Element) *List {
	if mark == nil || other == nil || mark.list != lst {
		return lst
	}
	for t := other.Front(); t != nil; t = t.Next() {
		mark = lst.insertValue(t.Value, mark)
	}
	return lst
}

// InsertAt inserts a new value at the position pos within the list and then return the list.
// If pos is either negative or more than the length of the list, the list is not modified.
func (lst *List) InsertAt(value interface{}, pos int) *List {
	if pos < 0 || pos >= lst.len {
		return lst
	}
	var elm *Element
	var i int
	for i, elm = 0, lst.Front(); (i+1 < pos) && (elm != nil); i, elm = i+1, elm.Next() {
	}
	if elm == nil {
		return lst
	}
	lst.InsertAfter(value, elm)
	return lst
}

// InsertListAt inserts another list at the position pos within the list and then return the list.
// If pos is either negative or more than the length of the list or the other list is nil, the list is not modified.
func (lst *List) InsertListAt(other *List, pos int) *List {
	if other == nil || pos < 0 || pos >= lst.len {
		return lst
	}
	if lst.len < pos {
		lst.PushBackList(other)
	} else {
		elm := lst.Nth(pos)
		lst.InsertListBefore(other, elm)
	}
	return lst
}

// InsertListBefore inserts a new list into the current list lst before mark and return the current list
// If either mark or the new list is nil or mark is not in the current list then the current list is not modified
func (lst *List) InsertListBefore(other *List, mark *Element) *List {
	if mark.list != lst || other == nil || mark == nil {
		return lst
	}
	for t := other.Back(); t != nil; t = t.Prev() {
		mark = lst.insertValue(t.Value, mark)
	}
	return lst
}

// MoveToFront moves element elm to the front of list lst.
// If elm is not an element of lst or nil, the list is not modified.
func (lst *List) MoveToFront(elm *Element) *List {
	if elm == nil || elm.list != lst || lst.root.next == elm {
		return lst
	}
	lst.move(elm, &lst.root)
	return lst
}

// MoveToBack moves element elm to the back of list lst.
// If elm is not an element of lst or nil, the list is not modified.
func (lst *List) MoveToBack(elm *Element) *List {
	if elm == nil || elm.list != lst || lst.root.prev == elm {
		return lst
	}
	lst.move(elm, lst.root.prev)
	return lst
}

// MoveBefore moves element elm to its new position before mark.
// If elm or mark is not an element of lst or elm == mark or either elm or mark is nil, the list is not modified.
func (lst *List) MoveBefore(elm, mark *Element) *List {
	if elm == nil || mark == nil || elm.list != lst || elm == mark || mark.list != lst {
		return lst
	}
	lst.move(elm, mark.prev)
	return lst
}

// MoveAfter moves element elm to its new position after mark.
// If elm or mark is not an element of lst or elm == mark or either elm or mark is nil, the list is not modified.
func (lst *List) MoveAfter(elm, mark *Element) *List {
	if elm == nil || mark == nil || elm.list != lst || elm == mark || mark.list != lst {
		return lst
	}
	lst.move(elm, mark)
	return lst
}

// PushBackList inserts a copy of a new list at the back of list lst.
// The lists lst and other may be the same. If other is nil nothing is changed.
func (lst *List) PushBackList(other *List) *List {
	if other == nil {
		return lst
	}
	lst.lazyInit()
	for i, elm := other.Len(), other.Front(); i > 0; i, elm = i-1, elm.Next() {
		lst.insertValue(elm.Value, lst.root.prev)
	}
	return lst
}

// PushFrontList inserts a copy of a new list at the front of list lst.
// The lists lst and other may be the same. If other is nil nothing is changed.
func (lst *List) PushFrontList(other *List) *List {
	if other == nil {
		return lst
	}
	lst.lazyInit()
	for i, elm := other.Len(), other.Back(); i > 0; i, elm = i-1, elm.Prev() {
		lst.insertValue(elm.Value, &lst.root)
	}
	return lst
}

// RemoveFunc will go through the list and based on the result of the application of the provided function
// it will remove the specific element from the list. The position in the list and the element at that position
// is provided to call the function which will return either true or false.
// What is left of the current list is returned.
func (lst *List) RemoveFunc(predFunc func(int, interface{}) bool) *List {
	i := 0
	elm := lst.Front()
	for elm != nil {
		nxt := elm.Next()
		if predFunc(i, elm.Value) {
			lst.Remove(elm)
		}
		elm = nxt
		i++
	}
	return lst
}

func (lst *List) PopFront() interface{} {
	if lst.Len() == 0 {
		return nil
	}
	tmp := lst.Front()
	lst.Remove(tmp)
	return tmp.Value
}

func (lst *List) PopBack() interface{} {
	if lst.Len() == 0 {
		return nil
	}
	tmp := lst.Back()
	lst.Remove(tmp)
	return tmp.Value
}
