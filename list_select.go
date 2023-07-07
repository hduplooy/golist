package list

// Len returns the number of elements of list lst.
func (lst *List) Len() int { return lst.len }

// Front returns the first element of list lst or nil if the list is empty.
func (lst *List) Front() *Element {
	if lst.len == 0 {
		return nil
	}
	return lst.root.next
}

// Back returns the last element of list lst or nil if the list is empty.
func (lst *List) Back() *Element {
	if lst.len == 0 {
		return nil
	}
	return lst.root.prev
}

// Nth returns the nth element in the list.
// If n is less than 0 then NthRev value is returned.
// If n is more than the length of the list nil is returned
func (lst *List) Nth(n int) *Element {
	if n < 0 {
		return lst.NthRev(-n)
	}
	if n >= lst.len {
		return nil
	}
	var elm *Element
	var i int
	for i, elm = 0, lst.Front(); (i < n) && (elm != nil); i, elm = i+1, elm.Next() {
	}
	return elm
}

// NthRev returns the nth last element in the list.
// If n is less than 0 then Nth value is returned.
// If n is more than the length of the list nil is returned
func (lst *List) NthRev(n int) *Element {
	if n < 0 {
		return lst.Nth(-n)
	}
	if n >= lst.len {
		return nil
	}
	var elm *Element
	var i int
	for i, elm = 0, lst.Back(); (i < n) && (elm != nil); i, elm = i+1, elm.Prev() {
	}
	return elm
}

// FirstN returns the first n elements of the list as a new list.
// If n is less than or equal to zero an empty list is returned.
func (lst *List) FirstN(n int) *List {
	result := New()
	if n <= 0 {
		return result
	}
	for i, elm := 0, lst.Front(); (i < n) && (elm != nil); i, elm = i+1, elm.Next() {
		result.PushBack(elm.Value)
	}
	return result
}

// LastN returns the last n elements of the list as a new list.
// If n is less than or equal to zero an empty list is returned.
func (lst *List) LastN(n int) *List {
	result := New()
	if n <= 0 {
		return result
	}
	for i, elm := 0, lst.Back(); (i < n) && (elm != nil); i, elm = i+1, elm.Prev() {
		result.PushFront(elm.Value)
	}
	return result
}

// SubList returns a sublist of the current list from the start position to the end position.
// If the start is negative or the end is before the start an empty list is returned.
func (lst *List) SubList(start, end int) *List {
	result := New()
	if end < start || start < 0 {
		return result
	}

	for i, elm := 0, lst.Front(); (i < end) && (elm != nil); i, elm = i+1, elm.Next() {
		if i >= start {
			result.PushBack(elm.Value)
		}
	}
	return result
}

func (lst *List) Rest() *List {
	result := New()
	if lst.Len() <= 1 {
		return result
	}
	for elm := lst.Front().Next(); elm != nil; elm = elm.Next() {
		result.PushBack(elm.Value)
	}

	return result
}

// Reverse returns a new list with the elements in the reverse order.
func (lst *List) Reverse() *List {
	result := New()
	for elm := lst.Back(); elm != nil; elm = elm.Prev() {
		result.PushBack(elm.Value)
	}
	return result
}
