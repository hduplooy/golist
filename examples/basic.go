package main

import (
	"fmt"
	list "github.com/hduplooy/list"
	"strconv"
	"strings"
)

func main() {
	names := []string{"hansel", "gretel", "grendel", "clunky 2", "don2key", "butt2head", "jughead", "simpleton", "bart", "calvin", "hobbs", "john", "peter"}
	list.ToList(names).Filter(func(pos int, val interface{}) bool {
		return !strings.HasPrefix(val.(string), "b")
	}).Map(func(pos int, val interface{}) interface{} {
		return strings.TrimSpace(strings.ReplaceAll(val.(string), "2", ""))
	}).Reverse().SubList(5, 8).ForEach(func(pos int, val interface{}) {
		fmt.Printf("%d. %s\n", pos, val.(string))
	})

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	ans := list.ToList(primes).Fold(0, func(val1, val2 interface{}) interface{} {
		return val1.(int) + val2.(int)
	}).(int)
	fmt.Printf("Sum is %d\n", ans)
	sq, err := list.ToList(primes).ToArray(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq.([]int))

	list.Map(func(vals []interface{}) interface{} {
		return strconv.Itoa(vals[0].(int)) + " " + vals[1].(string)
	}, list.ToList(primes[0:5]), list.ToList(names[0:5])).ForEach(func(pos int, val interface{}) {
		fmt.Printf("%d. %s\n", pos, val.(string))
	})

	fmt.Println("---------------------")
	thelist := list.ToList(names)
	val := thelist.Nth(3)
	fmt.Println(val.Value)

	val = thelist.NthRev(4)
	fmt.Println(val.Value)

	fmt.Println("---------------------")
	thelist.InsertAt("ZZZZ", 2).ForEach(func(pos int, val interface{}) {
		fmt.Printf("%d. %s\n", pos, val.(string))
	})
	fmt.Println("---------------------")

	thelist.RemoveFunc(func(pos int, v interface{}) bool {
		return v.(string) > "d"
	}).ForEach(func(pos int, val interface{}) {
		fmt.Printf("%d. %s\n", pos, val.(string))
	})
	fmt.Println("---------------------")

}
