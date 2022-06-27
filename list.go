// My version with some stuff added

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
package list

import (
	"fmt"
	"reflect"
)

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove removes e from its list, decrements l.len
func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

// move moves e to next to at.
func (l *List) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

func (l *List) InsertListAfter(other *List, mark *Element) {
	for t := other.Front(); t != nil; t = t.Next() {
		mark = l.InsertAfter(t.Value, mark)
	}
}

func (l *List) InsertListBefore(other *List, mark *Element) {
	for t := other.Back(); t != nil; t = t.Prev() {
		mark = l.InsertBefore(t.Value, mark)
	}
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

func (l *List) FirstN(n int) *List {
	result := New()
	for i, elm := 0, l.Front(); (i < n) && (elm != nil); i, elm = i+1, elm.Next() {
		result.PushBack(elm.Value)
	}
	return result
}

func (l *List) LastN(n int) *List {
	result := New()
	for i, elm := 0, l.Back(); (i < n) && (elm != nil); i, elm = i+1, elm.Prev() {
		result.PushFront(elm.Value)
	}
	return result
}

func (l *List) SubList(strt, end int) *List {
	result := New()
	for i, elm := 0, l.Front(); (i < end) && (elm != nil); i, elm = i+1, elm.Next() {
		if i >= strt {
			result.PushBack(elm.Value)
		}
	}
	return result
}

func (l *List) Filter(pr func(interface{}) bool) *List {
	other := New()
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		if pr(elm.Value) {
			other.PushBack(elm.Value)
		}
	}
	return other
}

func (l *List) Map(pr func(interface{}) interface{}) *List {
	other := New()
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		other.PushBack(pr(elm.Value))
	}
	return other
}

func (l *List) ForEach(pr func(interface{})) {
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		pr(elm.Value)
	}
}

func ToList(lst interface{}) *List {
	if reflect.TypeOf(lst).Kind() != reflect.Slice {
		return nil
	}
	list := New()
	for i := 0; i < reflect.ValueOf(lst).Len(); i++ {
		val := reflect.ValueOf(lst).Index(i).Interface()
		list.PushBack(val)
	}
	return list
}

func (l *List) Count(pr func(interface{}) bool) int {
	cnt := 0
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		if pr(elm.Value) {
			cnt++
		}
	}
	return cnt
}

func (l *List) DeMux(pr func(interface{}) string) map[string]*List {
	result := make(map[string]*List)

	for elm := l.Front(); elm != nil; elm = elm.Next() {
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

func (l *List) Fold(init interface{}, f func(val1, val2 interface{}) interface{}) interface{} {
	if l.Len() < 2 {
		return nil
	}
	ans := init
	for elm := l.Front(); elm != nil; elm = elm.Next() {
		ans = f(ans, elm.Value)
	}

	return ans
}

func (l *List) Reverse() *List {
	result := New()
	for elm := l.Back(); elm != nil; elm = elm.Prev() {
		result.PushBack(elm.Value)
	}
	return result
}

func (l *List) ToArray(dstif interface{}) (interface{}, error) {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(dstif)), l.Len(), l.Len())
	for i, elm := 0, l.Front(); elm != nil; i, elm = i+1, elm.Next() {
		val := elm.Value
		if val != nil && reflect.TypeOf(val) != reflect.TypeOf(dstif) {
			return nil, fmt.Errorf("Types of list not same as destination type")
		}
		slice.Index(i).Set(reflect.ValueOf(val))
	}

	return slice.Interface(), nil
}

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

//names := []string{"hansie","jorsie #","donsie #","kloekie","donkie","dofkop","jughead","simpleton","klipkop","pateet","noone","john","peter"}
//list.ToList(names).Filter(func(val interface{}) bool {
//	return strings.Index(val.(string),"#") < 0
//}).Map(func(val interface{}) interface{} {
//return strings.TrimSpace(strings.ReplaceAll(val.(string), "2",""))
//}).Reverse().SubList(5,8).ForEach(func(val interface{}) {
//	fmt.Println(val.(string))
//})
//
//primes := []int{2,3,5,7,11,13,17,19,23,29,31,37}
//ans := list.ToList(primes).Fold(0,func(val1,val2 interface{}) interface{} {
//return val1.(int) + val2.(int)
//}).(int)
//fmt.Printf("Sum is %d\n",ans)
//sq, err := list.ToList(primes).ToArray(3)
//if err != nil {
//fmt.Println(err)
//}
//fmt.Println(sq.([]int))
//
//list.Map(func(vals []interface{}) interface{} {
//return strconv.Itoa(vals[0].(int)) + " " + vals[1].(string)
//},list.ToList(primes[0:5]),list.ToList(names[0:5])).ForEach(func(val interface{}) {
//	fmt.Println(val.(string))
//})
