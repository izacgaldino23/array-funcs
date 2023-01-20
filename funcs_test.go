package arrayfuncs_test

import (
	"testing"

	arrayFuncs "github.com/izacgaldino23/array-funcs"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {

	t.Run("TestAnyToArrayKind", func(t *testing.T) {
		originalKind := []int{1, 2, 3, 4}

		new := arrayFuncs.AnyToArrayKind(originalKind)

		for i := range originalKind {
			assert.Equal(t, originalKind[i], new[i])
		}
	})

	t.Run("TestToOriginalKind", func(t *testing.T) {
		// Array create
		a := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		originalKind := a.ToOriginalKind()

		for i := range a {
			assert.Equal(t, a[i], originalKind[i])
		}
	})

	t.Run("TestAt", func(t *testing.T) {
		// Array create
		a := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		t.Run("PositiveValid", func(t *testing.T) {
			index := 3
			res := a.At(index)

			assert.Equal(t, *res, a[index])
		})

		t.Run("PositiveInvalid", func(t *testing.T) {
			index := len(a) + 1
			res := a.At(index)

			assert.Nil(t, res)
		})

		t.Run("Negative", func(t *testing.T) {
			index := -1
			res := a.At(index)

			assert.Equal(t, *res, a[len(a)+index])
		})
	})

	t.Run("TestConcat", func(t *testing.T) {
		var (
			// Array create
			a = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
			b = arrayFuncs.Array[int]{6, 7, 8, 9, 10}
		)

		c := a.Concat(&b)

		// Test new length
		assert.Equal(t, len(c), len(a)+len(b))

		// Test elements
		for i := range a {
			assert.Equal(t, a[i], c[i])
		}

		for i := range b {
			assert.Equal(t, b[i], c[i+len(b)])
		}
	})

	t.Run("TestEvery", func(t *testing.T) {
		// Array create
		a := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		pass := func(v *int, index int) bool {
			return *v < 6
		}

		dontPass := func(v *int, index int) bool {
			return *v < 3
		}

		// All results are under 6
		assert.True(t, a.Every(pass))

		// Some values are greater then 3
		assert.False(t, a.Every(dontPass))
	})

	t.Run("TestFill", func(t *testing.T) {
		var (
			a       = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
			b       = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
			result1 = []int{1, 2, 10, 10, 10}
			result2 = []int{1, 2, 10, 10, 5}
		)

		// Test without end
		a.Fill(10, 2)
		for i := range a {
			assert.Equal(t, a[i], result1[i])
		}

		// Test with end
		b.Fill(10, 2, 3)
		for i := range b {
			assert.Equal(t, b[i], result2[i])
		}
	})

	t.Run("TestFilter", func(t *testing.T) {
		// Array create
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		res := s.Filter(func(v *int, i int) bool {
			return *v > 3
		})

		assert.Equal(t, len(res), 2)

		assert.Equal(t, res[0], 4)

		assert.Equal(t, res[1], 5)
	})

	t.Run("TestFind", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		find := func(v *int, i int) bool {
			return *v == 3
		}

		assert.Equal(t, *s.Find(find), 3)

		notFind := func(v *int, i int) bool {
			return *v == 10
		}

		// Not found
		assert.Nil(t, s.Find(notFind))
	})

	t.Run("TestFindIndex", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		find := func(v *int, i int) bool {
			return *v == 3
		}

		notFind := func(v *int, i int) bool {
			return *v == 10
		}

		// Find
		assert.Equal(t, *s.FindIndex(find), 2)

		// Not Found
		assert.Nil(t, s.FindIndex(notFind))
	})

	t.Run("TestFindLast", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 3, 5}

		find := func(v *int, i int) bool {
			return *v == 3
		}

		assert.Equal(t, *s.FindLast(find), 3)

		notFind := func(v *int, i int) bool {
			return *v == 10
		}

		// Not found
		assert.Nil(t, s.Find(notFind))
	})

	t.Run("TestFindLastIndex", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 3, 5}

		find := func(v *int, i int) bool {
			return *v == 3
		}

		notFind := func(v *int, i int) bool {
			return *v == 10
		}

		// Find
		assert.Equal(t, *s.FindLastIndex(find), 3)

		// Not Found
		assert.Nil(t, s.FindLastIndex(notFind))
	})

	t.Run("TestForEach", func(t *testing.T) {
		var (
			a     = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
			sum   = 0
			total = 15
		)

		// Basic sum and compare result
		a.ForEach(func(value, index int, array *[]int) {
			sum += value
		})

		assert.Equal(t, sum, total)

		// Alter the original Array
		a.ForEach(func(value, index int, array *[]int) {
			*array = append(*array, value+1)
		})

		// Test if the elementes count is changed to double
		assert.Equal(t, len(a), 10)
	})

	t.Run("TestGroup", func(t *testing.T) {
		var (
			a = arrayFuncs.Array[int]{0, 1, 2, 3, 4, 5}
		)

		group := a.Group(func(value, index int) any {
			if value == 0 {
				return nil
			}

			kind := "odd"

			if value%2 == 0 {
				kind = "even"
			}

			return kind
		})

		for i := range group {
			for j := range group[i] {
				if i == "odd" {
					assert.Equal(t, group[i][j]%2, 1)
				} else {
					assert.Equal(t, group[i][j]%2, 0)
				}
			}
		}
	})

	t.Run("TestIncludes", func(t *testing.T) {
		var (
			a = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		)

		// Include
		assert.True(t, a.Includes(5))

		// Doesn't include
		assert.False(t, a.Includes(0))
	})

	t.Run("TestIndexOf", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		// Find
		assert.Equal(t, s.IndexOf(2), 1)

		// Not Found
		assert.Equal(t, s.IndexOf(6), -1)
	})

	t.Run("TestJoin", func(t *testing.T) {
		var (
			a = arrayFuncs.Array[int]{1, 2, 3, 4, 5}
			b = arrayFuncs.Array[bool]{true, false, true}
			c = arrayFuncs.Array[float32]{10.5, 3.4}
			d = arrayFuncs.Array[Temp]{
				{"hello"},
				{"world"},
			}
			separator = "-"
		)

		assert.Equal(t, "1-2-3-4-5", a.Join(separator))
		assert.Equal(t, "true-false-true", b.Join(separator))
		assert.Equal(t, "10.5-3.4", c.Join(separator))
		assert.Equal(t, "hello-world", d.Join(separator))
	})

	t.Run("TestKeys", func(t *testing.T) {
		a := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		expected := []int{0, 1, 2, 3, 4}
		result := a.Keys()

		for i := range result {
			assert.Equal(t, expected[i], result[i])
		}
	})

	t.Run("TestLastIndexOf", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 2, 5}

		// Find
		assert.Equal(t, s.LastIndexOf(2), 3)

		// Not Found
		assert.Equal(t, s.LastIndexOf(6), -1)
	})

	t.Run("TestMap", func(t *testing.T) {
		// criação do slice
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		// aplicação do map
		s.Map(func(v *int, i int) {
			*v = *v * 2
		})
		// verificação se o map foi aplicado corretamente
		for i, v := range s {
			assert.Equal(t, (i+1)*2, v)
		}
	})

	t.Run("TestPop", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		t.Run("ValidArray", func(t *testing.T) {
			// Remove last element
			removed := s.Pop()

			// Check new length
			assert.Equal(t, 4, len(s))

			// Check the removed element
			assert.NotNil(t, removed)
			assert.Equal(t, 5, *removed)
		})

		t.Run("EmptyArray", func(t *testing.T) {
			// Validate empty Array
			s = arrayFuncs.Array[int]{}
			removed := s.Pop()

			// Check new length
			assert.Equal(t, 0, len(s))

			// Check the removed element
			assert.Nil(t, removed)
		})
	})

	t.Run("TestPush", func(t *testing.T) {
		// Array create
		a := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		a.Push(6, 7)

		assert.Equal(t, 7, len(a))
	})

	t.Run("TestReduce", func(t *testing.T) {
		t.Run("SimpleValues", func(t *testing.T) {
			s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

			sumFunction := func(accumulator any, currentValue, currentIndex int) any {
				return accumulator.(int) + currentValue
			}

			result := s.Reduce(sumFunction, 0)

			assert.Equal(t, 15, result)
		})

		t.Run("StructTest", func(t *testing.T) {
			s := arrayFuncs.Array[Temp]{
				{"Hello"},
				{"World"},
				{"Mother"},
				{"Fuckers"},
			}

			concatFunction := func(accumulator any, currentValue Temp, currentIndex int) any {
				return accumulator.(string) + " " + currentValue.msg
			}

			result := s.Reduce(concatFunction, "->")

			assert.Equal(t, "-> Hello World Mother Fuckers", result)
		})
	})

	t.Run("TestReduceRight", func(t *testing.T) {
		t.Run("SimpleValues", func(t *testing.T) {
			s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

			sumFunction := func(accumulator any, currentValue, currentIndex int) any {
				return accumulator.(int) + currentValue
			}

			result := s.ReduceRight(sumFunction, 0)

			assert.Equal(t, 15, result)
		})

		t.Run("StructTest", func(t *testing.T) {
			s := arrayFuncs.Array[Temp]{
				{"Fuckers"},
				{"Mother"},
				{"World"},
				{"Hello"},
			}

			concatFunction := func(accumulator any, currentValue Temp, currentIndex int) any {
				return accumulator.(string) + " " + currentValue.msg
			}

			result := s.ReduceRight(concatFunction, "->")

			assert.Equal(t, "-> Hello World Mother Fuckers", result)
		})
	})

	t.Run("TestReverse", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		expected := arrayFuncs.Array[int]{5, 4, 3, 2, 1}

		s.Reverse()

		for i := range expected {
			assert.Equal(t, expected[i], s[i])
		}
	})

	t.Run("TestShift", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

		t.Run("ValidArray", func(t *testing.T) {
			// Remove last element
			removed := s.Shift()

			// Check new length
			assert.Equal(t, 4, len(s))

			// Check the removed element
			assert.NotNil(t, removed)
			assert.Equal(t, 1, *removed)
		})

		t.Run("EmptyArray", func(t *testing.T) {
			// Validate empty Array
			s = arrayFuncs.Array[int]{}
			removed := s.Pop()

			// Check new length
			assert.Equal(t, 0, len(s))

			// Check the removed element
			assert.Nil(t, removed)
		})
	})

	t.Run("model", func(t *testing.T) {
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		_ = s
	})
}
