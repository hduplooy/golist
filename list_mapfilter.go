package list

// Filter applies the provided function (which returns a boolean) to return only those elements as a new list for which the function returns a true value.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) Filter(pr func(int, interface{}) bool) *List {
	other := New()
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		if pr(i, elm.Value) {
			other.PushBack(elm.Value)
		}
	}
	return other
}

// FilterTo applies the provided function (which returns a boolean) to return all elements up to this first one that the test returns true else everything is returned.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) FilterTo(pr func(int, interface{}) bool) *List {
	other := New()
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		other.PushBack(elm.Value)
		if pr(i, elm.Value) {
			break
		}
	}
	return other
}

// FilterFrom applies the provided function (which returns a boolean) and then return every element from there onwards.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) FilterFrom(pr func(int, interface{}) bool) *List {
	other := New()
	found := false
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		if pr(i, elm.Value) {
			found = true
		}
		if found {
			other.PushBack(elm.Value)
		}
	}
	return other
}

// Map applies the provided function on the list and form a new list from the returned values.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) Map(pr func(int, interface{}) interface{}) *List {
	other := New()
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		other.PushBack(pr(i, elm.Value))
	}
	return other
}

// ForEach just performs the provided function on each element.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) ForEach(pr func(int, interface{})) {
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		pr(i, elm.Value)
	}
}

// Count will count the number of elements for which the provided function is true.
// For each element the position in the list and the actual element is provided to the function.
func (lst *List) Count(pr func(int, interface{}) bool) int {
	cnt := 0
	for i, elm := 0, lst.Front(); elm != nil; i, elm = i+1, elm.Next() {
		if pr(i, elm.Value) {
			cnt++
		}
	}
	return cnt
}

// DeMux will apply the provided function and split the elements as a map based on the returned string value of the function.
func (lst *List) DeMux(pr func(interface{}) string) map[string]*List {
	result := make(map[string]*List)

	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		key := pr(elm.Value)
		tmp, ok := result[key]
		if !ok {
			tmp = New()
			result[key] = tmp
		}
		tmp.PushBack(elm.Value)
	}
	return result
}

// Fold will apply the provided function on the init value and the first of the list and then again on each of the rest of the list returning the last value obtained.
func (lst *List) Fold(init interface{}, f func(val1, val2 interface{}) interface{}) interface{} {
	if lst.Len() < 2 {
		return nil
	}
	ans := init
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		ans = f(ans, elm.Value)
	}

	return ans
}
