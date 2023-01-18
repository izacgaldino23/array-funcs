package arrayfuncs_test

import (
	"testing"

	arrayFuncs "github.com/izacgaldino23/array-funcs"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {

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

	t.Run("model", func(t *testing.T) {})

	t.Run("TestMap", func(t *testing.T) {
		// criação do slice
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		// aplicação do map
		s.Map(func(i int, v *int) {
			*v = *v * 2
		})
		// verificação se o map foi aplicado corretamente
		for i, v := range s {
			assert.Equal(t, (i+1)*2, v)
		}
	})

	t.Run("TestFilter", func(t *testing.T) {
		// criação do slice
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		// aplicação do filter
		res := s.Filter(func(v *int, i int) bool {
			return *v > 3
		})
		// verificação se o filter foi aplicado corretamente
		assert.Equal(t, len(res), 2)

		assert.Equal(t, res[0], 4)

		assert.Equal(t, res[1], 5)
	})

	t.Run("TestFind", func(t *testing.T) {
		// criação do slice
		s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}
		// aplicação do find
		res := s.Find(func(v *int, i int) bool {
			return *v == 3
		})
		// verificação se o find foi aplicado corretamente
		assert.Equal(t, res, 3)
	})
}
