package gset

import "sync"

type threadSafeSet struct {
	s threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{s: newThreadUnsafeSet()}
}

func (set *threadSafeSet) Size() int {
	set.RLock()
	defer set.RUnlock()
	l := len(set.s)
	return l
}

func (set *threadSafeSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *threadSafeSet) Contains(i ...interface{}) bool {
	set.RLock()
	defer set.RUnlock()

	ret := set.s.Contains(i...)
	return ret
}

func (set *threadSafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	set.RLock()
	defer set.RUnlock()

	for elem := range set.s {
		keys = append(keys, elem)
	}
	return keys
}

func (set *threadSafeSet) Add(i interface{}) bool {
	set.Lock()
	defer set.Unlock()
	ret := set.s.Add(i)
	return ret
}

func (set *threadSafeSet) Remove(i interface{}) {
	set.Lock()
	defer set.Unlock()
	delete(set.s, i)
}

func (set *threadSafeSet) Clear() {
	set.Lock()
	defer set.Unlock()
	set.s = newThreadUnsafeSet()
}

func (set *threadSafeSet) String() string {
	set.RLock()
	defer set.RUnlock()

	ret := set.s.String()
	return ret
}
