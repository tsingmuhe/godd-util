package gset

type Set interface {
	Size() int

	IsEmpty() bool

	Contains(i ...interface{}) bool

	ToSlice() []interface{}

	Add(i interface{}) bool

	Remove(i interface{})

	//containsAll(other Set) bool
	//
	//addAll(other Set) bool
	//
	//retainAll(other Set) bool
	//
	//removeAll(other Set) bool

	Clear()

	String() string
}

func NewSet() Set {
	set := newThreadUnsafeSet()
	return &set
}

func NewSetFromSlice(s []interface{}) Set {
	a := NewSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}

func NewConcurrentSet() Set {
	set := newThreadSafeSet()
	return &set
}

func NewConcurrentSetFromSlice(s []interface{}) Set {
	set := newThreadSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}
