package gset

import (
	"fmt"
	"strings"
)

type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (set *threadUnsafeSet) Size() int {
	return len(*set)
}

func (set *threadUnsafeSet) IsEmpty() bool {
	return len(*set) == 0
}

func (set *threadUnsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}

	return true
}

func (set *threadUnsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	for elem := range *set {
		keys = append(keys, elem)
	}

	return keys
}

func (set *threadUnsafeSet) Add(i interface{}) bool {
	_, found := (*set)[i]
	if found {
		return false //False if it existed already
	}

	(*set)[i] = struct{}{}
	return true
}

func (set *threadUnsafeSet) Remove(i interface{}) {
	delete(*set, i)
}

func (set *threadUnsafeSet) Clear() {
	*set = newThreadUnsafeSet()
}

func (set *threadUnsafeSet) String() string {
	items := make([]string, 0, len(*set))

	for elem := range *set {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
