package list

import (
	"fmt"
	"reflect"
)

// ToList will take an array and make a list out of it
func ToList(array interface{}) *List {
	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return nil
	}
	list := New()
	for i := 0; i < reflect.ValueOf(array).Len(); i++ {
		val := reflect.ValueOf(array).Index(i).Interface()
		list.PushBack(val)
	}

	return list
}

// ToArray will take the current list and covert it to an array of the type specified by dstif
// If any of the elements are not of the type provided an error is returned
func (l *List) ToArray(destIF interface{}) (interface{}, error) {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(destIF)), l.Len(), l.Len())
	for i, elm := 0, l.Front(); elm != nil; i, elm = i+1, elm.Next() {
		val := elm.Value
		if val != nil && reflect.TypeOf(val) != reflect.TypeOf(destIF) {
			return nil, fmt.Errorf("Types of list not same as destination type")
		}
		slice.Index(i).Set(reflect.ValueOf(val))
	}

	return slice.Interface(), nil
}
