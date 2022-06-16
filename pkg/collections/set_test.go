package collections

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Empty()

func TestEmpty(t *testing.T) {
	s1 := NewSet(1)
	assert.False(t, s1.Empty())
	s2 := NewSet[int]()
	assert.True(t, s2.Empty())
}

// Len()

func TestLen(t *testing.T) {
	s1 := NewSet[int]()
	assert.Equal(t, 0, s1.Len())
	s2 := NewSet(1)
	assert.Equal(t, 1, s2.Len())
	s3 := NewSet(1, 2)
	assert.Equal(t, 2, s3.Len())
}

// List()

func TestListInt(t *testing.T) {
	s := NewSet(2, 3, 1)
	assert.Equal(t, []int{1, 2, 3}, s.List())
}

func TestListString(t *testing.T) {
	s := NewSet("b", "a", "c")
	assert.Equal(t, []string{"a", "b", "c"}, s.List())
}

// Has(v T)

func TestHas(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
	assert.True(t, s.Has(3))
	assert.False(t, s.Has(4))
}

// HasAll()

func TestHasAll(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.HasAll(1))
	assert.True(t, s.HasAll(2))
	assert.True(t, s.HasAll(3))
	assert.True(t, s.HasAll(1, 2))
	assert.True(t, s.HasAll(2, 3))
	assert.True(t, s.HasAll(1, 3))
	assert.True(t, s.HasAll(1, 2, 3))
	assert.False(t, s.HasAll(4))
	assert.False(t, s.HasAll(3, 4))
	assert.False(t, s.HasAll(1, 2, 3, 4))
	assert.False(t, s.HasAll())
}

// HasAllFromSet()

func TestHasAllFromSet(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.HasAllFromSet(NewSet(1)))
	assert.True(t, s.HasAllFromSet(NewSet(2)))
	assert.True(t, s.HasAllFromSet(NewSet(3)))
	assert.True(t, s.HasAllFromSet(NewSet(1, 2)))
	assert.True(t, s.HasAllFromSet(NewSet(2, 3)))
	assert.True(t, s.HasAllFromSet(NewSet(1, 3)))
	assert.True(t, s.HasAllFromSet(NewSet(1, 2, 3)))
	assert.False(t, s.HasAllFromSet(NewSet(4)))
	assert.False(t, s.HasAllFromSet(NewSet(3, 4)))
	assert.False(t, s.HasAllFromSet(NewSet(1, 2, 3, 4)))
	assert.False(t, s.HasAllFromSet(NewSet[int]()))
}

// HasAny()

func TestHasAny(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.HasAny(1))
	assert.True(t, s.HasAny(2))
	assert.True(t, s.HasAny(3))
	assert.True(t, s.HasAny(1, 2))
	assert.True(t, s.HasAny(2, 3))
	assert.True(t, s.HasAny(1, 3))
	assert.True(t, s.HasAny(1, 2, 3))
	assert.True(t, s.HasAny(3, 4))
	assert.True(t, s.HasAny(1, 2, 3, 4))
	assert.False(t, s.HasAny(4))
	assert.False(t, s.HasAny(4, 5))
	assert.False(t, s.HasAny())
}

// HasAnyFromSet()

func TestHasAnyFromSet(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.HasAnyFromSet(NewSet(1)))
	assert.True(t, s.HasAnyFromSet(NewSet(2)))
	assert.True(t, s.HasAnyFromSet(NewSet(3)))
	assert.True(t, s.HasAnyFromSet(NewSet(1, 2)))
	assert.True(t, s.HasAnyFromSet(NewSet(2, 3)))
	assert.True(t, s.HasAnyFromSet(NewSet(1, 3)))
	assert.True(t, s.HasAnyFromSet(NewSet(1, 2, 3)))
	assert.True(t, s.HasAnyFromSet(NewSet(3, 4)))
	assert.True(t, s.HasAnyFromSet(NewSet(1, 2, 3, 4)))
	assert.False(t, s.HasAnyFromSet(NewSet(4)))
	assert.False(t, s.HasAnyFromSet(NewSet(4, 5)))
	assert.False(t, s.HasAnyFromSet(NewSet[int]()))
}

// Union()

func TestUnion(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			var s Set[int]
			s = NewSet(s1...).Union(NewSet(s2...))
			assert.Equal(t, sExpected, s.List())
			s = NewSet(s2...).Union(NewSet(s1...))
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{1, 2, 3, 4, 5, 6})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{1, 2, 3, 4})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{1, 2, 3})
	f("two empty sets", []int{}, []int{}, []int{})
}

// Intersection()

func TestIntersection(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			var s Set[int]
			s = NewSet(s1...).Intersection(NewSet(s2...))
			assert.Equal(t, sExpected, s.List())
			s = NewSet(s2...).Intersection(NewSet(s1...))
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{2, 3})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{})
	f("two empty sets", []int{}, []int{}, []int{})
}

// Difference()

func TestDifference(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			s := NewSet(s1...).Difference(NewSet(s2...))
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{1, 2, 3})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{1})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{1, 2, 3})
	f("from empty set", []int{}, []int{2, 1, 3}, []int{})
	f("two empty sets", []int{}, []int{}, []int{})
}

// SymmetricDifference()

func TestSymmetricDifference(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			var s Set[int]
			s = NewSet(s1...).SymmetricDifference(NewSet(s2...))
			assert.Equal(t, sExpected, s.List())
			s = NewSet(s2...).SymmetricDifference(NewSet(s1...))
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{1, 2, 3, 4, 5, 6})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{1, 4})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{1, 2, 3})
	f("from empty set", []int{}, []int{2, 1, 3}, []int{1, 2, 3})
	f("two empty sets", []int{}, []int{}, []int{})
}

// Add()

func TestAdd(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			s := NewSet(s1...).Add(s2...)
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{1, 2, 3, 4, 5, 6})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{1, 2, 3, 4})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{1, 2, 3})
	f("two empty sets", []int{}, []int{}, []int{})
}

// Remove()

func TestRemove(t *testing.T) {
	f := func(description string, s1, s2, sExpected []int) {
		t.Run(description, func(t *testing.T) {
			s := NewSet(s1...).Remove(s2...)
			assert.Equal(t, sExpected, s.List())
		})
	}

	f("without intersection", []int{2, 1, 3}, []int{6, 5, 4}, []int{1, 2, 3})
	f("with intersection", []int{2, 1, 3}, []int{4, 2, 3}, []int{1})
	f("with empty set", []int{2, 1, 3}, []int{}, []int{1, 2, 3})
	f("from empty set", []int{}, []int{2, 1, 3}, []int{})
	f("two empty sets", []int{}, []int{}, []int{})
}

// IsSubsetOf()

func TestIsSubsetOf(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.IsSubsetOf(NewSet(2, 1, 3)))
	assert.True(t, s.IsSubsetOf(NewSet(1, 2, 3, 4)))
	assert.True(t, NewSet[int]().IsSubsetOf(s))
	assert.True(t, NewSet[int]().IsSubsetOf(NewSet[int]()))
	assert.False(t, s.IsSubsetOf(NewSet(1, 2)))
	assert.False(t, s.IsSubsetOf(NewSet[int]()))
}

// IsSupersetOf()

func TestIsSupersetOf(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.True(t, s.IsSupersetOf(NewSet(2, 1, 3)))
	assert.True(t, s.IsSupersetOf(NewSet(1, 2)))
	assert.True(t, s.IsSupersetOf(NewSet[int]()))
	assert.True(t, NewSet[int]().IsSupersetOf(NewSet[int]()))
	assert.False(t, NewSet[int]().IsSupersetOf(s))
	assert.False(t, s.IsSupersetOf(NewSet(1, 2, 3, 4)))
}

// Equal()

func TestEqual(t *testing.T) {
	assert.True(t, NewSet(1, 2, 3).Equal(NewSet(2, 1, 3)))
	assert.True(t, NewSet[int]().Equal(NewSet[int]()))
	assert.False(t, NewSet(1, 2, 3).Equal(NewSet(1, 2, 3, 4)))
	assert.False(t, NewSet(1, 2, 3, 4).Equal(NewSet(1, 2, 3)))
	assert.False(t, NewSet(1, 2, 3).Equal(NewSet[int]()))
	assert.False(t, NewSet[int]().Equal(NewSet(1, 2, 3)))
}

// SetMap()

func TestSetMap(t *testing.T) {
	s := NewSet(1, 2, 3)
	sExpected := NewSet("1", "2", "3")
	sResult := SetMap(s, func(i int) string { return fmt.Sprint(i) })
	assert.True(t, sExpected.Equal(sResult))
}

func TestSetMapWithIntersection(t *testing.T) {
	s := NewSet(-3, -2, -1, 0, 1, 2, 3)
	sExpected := NewSet("0", "1", "2", "3")
	sResult := SetMap(s, func(i int) string {
		if i < 0 {
			i = -i
		}
		return fmt.Sprint(i)
	})
	assert.True(t, sExpected.Equal(sResult))
}

// SetMapErr()

func TestSetMapErr(t *testing.T) {
	s := NewSet(1, 2, 3)
	sExpected := NewSet("1", "2", "3")
	sResult, err := SetMapErr(s, func(i int) (string, error) { return fmt.Sprint(i), nil })
	require.NoError(t, err)
	assert.True(t, sExpected.Equal(sResult))
}

func TestSetMapErrWithIntersection(t *testing.T) {
	s := NewSet(-3, -2, -1, 0, 1, 2, 3)
	sExpected := NewSet("0", "1", "2", "3")
	sResult, err := SetMapErr(s, func(i int) (string, error) {
		if i < 0 {
			i = -i
		}
		return fmt.Sprint(i), nil
	})
	require.NoError(t, err)
	assert.True(t, sExpected.Equal(sResult))
}

func TestSetMapErrInterrupted(t *testing.T) {
	s := NewSet(1, 2, 3)
	_, err := SetMapErr(s, func(i int) (string, error) {
		if i == 2 {
			return "", errors.New("")
		}
		return fmt.Sprint(i), nil
	})
	require.Error(t, err)
}

// SetFilter()

func TestSetFilter(t *testing.T) {
	s := NewSet(-3, -2, -1, 0, 1, 2, 3)
	sExpected := NewSet(1, 2, 3)
	sResult := SetFilter[int](s, func(i int) bool { return i > 0 })
	assert.True(t, sExpected.Equal(sResult))
}

// SetFilterErr()

func TestSetFilterErr(t *testing.T) {
	s := NewSet(-3, -2, -1, 0, 1, 2, 3)
	sExpected := NewSet(1, 2, 3)
	sResult, err := SetFilterErr[int](s, func(i int) (bool, error) { return i > 0, nil })
	require.NoError(t, err)
	assert.True(t, sExpected.Equal(sResult))
}

func TestSetFilterErrInterrupted(t *testing.T) {
	s := NewSet(-3, -2, -1, 0, 1, 2, 3)
	_, err := SetFilterErr[int](s, func(i int) (bool, error) {
		if i == 1 {
			return false, errors.New("")
		}
		return i > 0, nil
	})
	require.Error(t, err)
}

// SetReduce()

func TestSetReduceSumInt(t *testing.T) {
	s := NewSet(1, 2, 3)
	sum := SetReduce(s, 5, AccumulatorNumberSum[int]())
	assert.Equal(t, 11, sum)
}

func TestSetReduceSumFloat32(t *testing.T) {
	s := NewSet[float32](1., 2., 3.)
	sum := SetReduce(s, 5., AccumulatorNumberSum[float32]())
	assert.Equal(t, float32(11.), sum)
}

func TestSetReduceSumKindOfFloat64(t *testing.T) {
	type testType float64
	s := NewSet[testType](1., 2., 3.)
	sum := SetReduce(s, 5., AccumulatorNumberSum[testType]())
	assert.Equal(t, testType(11.), sum)
}

func TestSetReduceMulKindOfInt64(t *testing.T) {
	type testType int64
	s := NewSet[testType](1, 2, 3)
	sum := SetReduce(s, -1, AccumulatorNumberMul[testType]())
	assert.Equal(t, testType(-6), sum)
}

func TestSetReduceConcatString(t *testing.T) {
	s := NewSet("a", "b", "c")
	result := SetReduce(s, "_", func(partialResult, element string) string { return partialResult + element })
	assert.Equal(t, "_abc", result)
}

// SetReduceErr()

func TestSetReduceErr(t *testing.T) {
	s := NewSet(1, 2, 3)
	sum, err := SetReduceErr(s, 5, func(partialResult, element int) (int, error) { return partialResult + element, nil })
	require.NoError(t, err)
	assert.Equal(t, 11, sum)
}

func TestSetReduceErrInterrupted(t *testing.T) {
	s := NewSet(1, 2, 3)
	_, err := SetReduceErr(s, 5, func(partialResult, element int) (int, error) {
		if element == 3 {
			return 0, errors.New("")
		}
		return partialResult + element, nil
	})
	require.Error(t, err)
}
