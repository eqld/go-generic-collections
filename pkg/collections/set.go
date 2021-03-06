package collections

import (
	"fmt"
	"sort"
)

// Set is immutable generic set of elements of comparable type.
type Set[T comparable] struct {
	list []T
	set  map[T]struct{}
}

// NewSet creates and returns new immutable set of given elements.
func NewSet[T comparable](values ...T) (set Set[T]) {

	set.set = make(map[T]struct{}, len(values))
	for _, v := range values {
		set.set[v] = struct{}{}
	}

	set.list = make([]T, 0, len(set.set))
	for v := range set.set {
		set.list = append(set.list, v)
	}

	sort.Slice(set.list, func(left, right int) bool { return fmt.Sprint(set.list[left]) < fmt.Sprint(set.list[right]) })
	return
}

// Empty returns `true` if the set is empty (i.e. has 0 elements).
func (s Set[T]) Empty() bool {
	return len(s.list) == 0
}

// Len returns number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.list)
}

// List returns ordered list of elements of the set. Order of elements in the list is stable.
func (s Set[T]) List() (list []T) {
	list = make([]T, len(s.list))
	copy(list, s.list)
	return list
}

// Has returns `true` if the set has given element.
func (s Set[T]) Has(v T) (ok bool) {
	_, ok = s.set[v]
	return
}

// HasAll returns `true` if the set has all of given elements. It returns `false` if provided list of elements is empty.
func (s Set[T]) HasAll(vv ...T) (ok bool) {
	for _, v := range vv {
		if _, ok = s.set[v]; !ok {
			break
		}
	}
	return
}

// HasAllFromSet returns `true` if the set has all elements of given set. It returns `false` if provided set is empty.
func (s Set[T]) HasAllFromSet(another Set[T]) (ok bool) {
	return s.HasAll(another.list...)
}

// HasAny returns `true` if the set has any of given elements. It returns `false` if provided list of elements is empty.
func (s Set[T]) HasAny(vv ...T) (ok bool) {
	for _, v := range vv {
		if _, ok = s.set[v]; ok {
			break
		}
	}
	return
}

// HasAnyFromSet returns `true` if the set has any elements of given set. It returns `false` if provided set is empty.
func (s Set[T]) HasAnyFromSet(another Set[T]) (ok bool) {
	return s.HasAny(another.list...)
}

// Union returns union of this set and another set.
func (s Set[T]) Union(another Set[T]) (result Set[T]) {
	result = NewSet(append(s.List(), another.list...)...)
	return
}

// Intersection returns intersection of this set and another set.
func (s Set[T]) Intersection(another Set[T]) (result Set[T]) {
	values := make([]T, 0)
	for _, v := range another.list {
		if s.Has(v) {
			values = append(values, v)
		}
	}
	result = NewSet(values...)
	return
}

// Difference returns difference of this set and another set.
func (s Set[T]) Difference(another Set[T]) (result Set[T]) {
	values := make([]T, 0)
	for _, v := range s.list {
		if !another.Has(v) {
			values = append(values, v)
		}
	}
	result = NewSet(values...)
	return
}

// SymmetricDifference returns symmetric difference of this set and another set.
func (s Set[T]) SymmetricDifference(another Set[T]) (result Set[T]) {
	result = s.Union(another).Difference(s.Intersection(another))
	return
}

// Add returns union of this set and given elements.
func (s Set[T]) Add(elements ...T) (result Set[T]) {
	result = s.Union(NewSet(elements...))
	return
}

// Remove returns difference of this set and given elements.
func (s Set[T]) Remove(elements ...T) (result Set[T]) {
	result = s.Difference(NewSet(elements...))
	return
}

// IsSubsetOf returns `true` if the set is a subset of another set.
func (s Set[T]) IsSubsetOf(another Set[T]) (ok bool) {
	if s.Empty() {
		// Empty set is a subset of any other set.
		ok = true
	} else {
		ok = another.HasAllFromSet(s)
	}
	return
}

// IsSupersetOf returns `true` if the set is a superset of another set.
func (s Set[T]) IsSupersetOf(another Set[T]) (ok bool) {
	ok = another.IsSubsetOf(s)
	return
}

// Equal returns `true` if sets are equal.
func (s Set[T]) Equal(another Set[T]) (ok bool) {
	if len(s.list) != len(another.list) {
		ok = false
	} else {
		ok = s.IsSubsetOf(another) && another.IsSubsetOf(s)
	}
	return
}

// SetMap maps each element from given set of T to R using provided mapper function and returns set of R.
func SetMap[T, R comparable](set Set[T], mapper func(T) R) (result Set[R]) {
	values := make([]R, 0, len(set.list))
	for _, v := range set.list {
		values = append(values, mapper(v))
	}
	result = NewSet(values...)
	return
}

// SetMapErr maps each element from given set of T to R using provided mapper function and returns set of R.
// It stops iterating over elements if given mapper function returns non-nil error.
func SetMapErr[T, R comparable](set Set[T], mapper func(T) (R, error)) (result Set[R], err error) {
	values := make([]R, 0, len(set.list))
	for _, v := range set.list {
		var r R
		if r, err = mapper(v); err != nil {
			return
		}
		values = append(values, r)
	}
	result = NewSet(values...)
	return
}

// SetFilter filters elements of given set using provided filter function and returns new set with filtered elements.
func SetFilter[T comparable](set Set[T], filter func(T) bool) (result Set[T]) {
	values := make([]T, 0)
	for _, v := range set.list {
		if filter(v) {
			values = append(values, v)
		}
	}
	result = NewSet(values...)
	return
}

// SetFilterErr filters elements of given set using provided filter function and returns new set with filtered elements.
// It stops iterating over elements if given filter function returns non-nil error.
func SetFilterErr[T comparable](set Set[T], filter func(T) (bool, error)) (result Set[T], err error) {
	values := make([]T, 0)
	for _, v := range set.list {
		var ok bool
		if ok, err = filter(v); err != nil {
			return
		} else if ok {
			values = append(values, v)
		}
	}
	result = NewSet(values...)
	return
}

// SetReduce reduces given set to single element using provided initial value and accumulator function.
func SetReduce[T comparable](set Set[T], initialValue T, accumulator func(partialResult, element T) T) (result T) {
	result = initialValue
	for _, v := range set.list {
		result = accumulator(result, v)
	}
	return
}

// SetReduceErr reduces given set to single element using provided initial value and accumulator function.
// It stops iterating over elements if given accumulator function returns non-nil error.
func SetReduceErr[T comparable](set Set[T], initialValue T, accumulator func(partialResult, element T) (T, error)) (result T, err error) {
	result = initialValue
	for _, v := range set.list {
		if result, err = accumulator(result, v); err != nil {
			return
		}
	}
	return
}
