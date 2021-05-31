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
	l := len(set.s)
	set.RUnlock()
	return l
}

func (set *threadSafeSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *threadSafeSet) Contains(i ...interface{}) bool {
	set.RLock()
	ret := set.s.Contains(i...)
	set.RUnlock()
	return ret
}

func (set *threadSafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	set.RLock()
	for elem := range set.s {
		keys = append(keys, elem)
	}
	set.RUnlock()
	return keys
}

func (set *threadSafeSet) Add(i interface{}) bool {
	set.Lock()
	ret := set.s.Add(i)
	set.Unlock()
	return ret
}

func (set *threadSafeSet) Remove(i interface{}) {
	set.Lock()
	delete(set.s, i)
	set.Unlock()
}

func (set *threadSafeSet) Clear() {
	set.Lock()
	set.s = newThreadUnsafeSet()
	set.Unlock()
}

func (set *threadSafeSet) String() string {
	set.RLock()
	ret := set.s.String()
	set.RUnlock()
	return ret
}
