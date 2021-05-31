package gset

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

const N = 1000

func Test_AddConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))

	for i := 0; i < len(ints); i++ {
		go func(i int) {
			s.Add(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, i := range ints {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

func Test_CardinalityConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		elems := s.Size()
		for i := 0; i < N; i++ {
			newElems := s.Size()
			if newElems < elems {
				t.Errorf("Cardinality shrunk from %v to %v", elems, newElems)
			}
		}
		wg.Done()
	}()

	for i := 0; i < N; i++ {
		s.Add(rand.Int())
	}

	wg.Wait()
}

func Test_ClearConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))

	for i := 0; i < len(ints); i++ {
		go func() {
			s.Clear()
			wg.Done()
		}()

		go func(i int) {
			s.Add(i)
		}(i)
	}

	wg.Wait()
}

func Test_ContainsConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)
	interfaces := make([]interface{}, 0)

	for _, v := range ints {
		s.Add(v)
		interfaces = append(interfaces, v)
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Contains(interfaces...)
			wg.Done()
		}()
	}

	wg.Wait()
}

func Test_RemoveConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for _, v := range ints {
		go func(i int) {
			s.Remove(i)
			wg.Done()
		}(v)
	}
	wg.Wait()

	if s.Size() != 0 {
		t.Errorf("Expected cardinality 0; got %v", s.Size())
	}
}

func Test_StringConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for range ints {
		go func() {
			_ = s.String()
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_ToSlice(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewConcurrentSet()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i int) {
			s.Add(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	setAsSlice := s.ToSlice()
	if len(setAsSlice) != s.Size() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

func Test_ToSliceDeadlock(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	set := NewConcurrentSet()
	workers := 10
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				set.Add(1)
				set.ToSlice()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
