package list

// Map is a utility function that will apply the provided function on the provided lists returning a new list based on the returns
func Map(f func([]interface{}) interface{}, lists ...*List) *List {
	result := New()
	curVals := make([]*Element, len(lists))
	done := false
	for i, lst := range lists {
		curVals[i] = lst.Front()
		if curVals[i] == nil {
			done = true
		}
	}
	for !done {
		vals := make([]interface{}, len(curVals))
		for i, elm := range curVals {
			vals[i] = elm.Value
			curVals[i] = curVals[i].Next()
			if curVals[i] == nil {
				done = true
			}
		}
		result.PushBack(f(vals))
	}
	return result
}
