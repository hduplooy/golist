# golist

Work in progress.

Extensions of the standard list functionality in go.

See the examples folder for examples on how to use this.

## Current list of functions.

**Back** returns the last element of list l or nil if the list is empty.

**Count** will count the number of elements for which the provided function is true

**DeMux** will apply the provided function and split the elements as a map based on the returned string value of the function

**Filter** applies the provided function (which returns a boolean) to return only those elements as a new list for which the function returns a true value

**FirstN** return the first n elements of the list as a new list

**Fold** will apply the provided function on the init value and the first of the list and then again on each of the rest of the list returning the last value obtained

**ForEach** just performs the provided function on each element

**Front** returns the first element of list l or nil if the list is empty.

**Init** initializes or clears list l.

**InsertAfter** inserts a new element e with value v immediately after mark and returns e.

**InsertBefore** inserts a new element e with value v immediately before mark and returns e.

**InsertListAfter** insters a new list into the current list *l* after mark and return the current list

**InsertListBefore** insters a new list into the current list *l* before mark and return the current list

**LastN** return the last n elements of the list as a new list

**Len** returns the number of elements of list l.

**Map** applies the provided function on the list and form a new list from the returned values

**Map** on its own is a utility function that will apply the provided function on the provided lists returning a new list based on the returns

**MoveAfter** moves element e to its new position after mark.

**MoveBefore** moves element e to its new position before mark.

**MoveToBack** moves element e to the back of list l.

**MoveToFront** moves element e to the front of list l.

**New** returns an initialized list.

**Next** returns the next list element or nil.

**Prev** returns the previous list element or nil.

**PushBack** inserts a new element e with value v at the back of list l and returns e.

**PushBackList** inserts a copy of another list at the back of list l.

**PushFront** inserts a new element e with value v at the front of list l and returns e.

**PushFrontList** inserts a copy of another list at the front of list l.

**Remove** removes e from l if e is an element of list l.

**Reverse** will return a new list with the elements in the reverse order

**SubList** returns a sublist of the current list from the strt position to the end position

**ToArray** will take the current list and covert it to an array of the type specified by dstif

**ToList** will take an array and make a list out of it

