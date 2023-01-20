package arrayfuncs

import (
	"strings"
)

// Array is a base array of any type that has many usable functions
type Array[T comparable] []T

/*
AnyToArrayKind receive a slice of T kind and return a Array[T]

	[]int will return Array[int]
*/
func AnyToArrayKind[T comparable](input []T) (res Array[T]) {
	res = make(Array[T], 0)

	for i := range input {
		res = append(res, input[i])
	}

	return
}

/*
ToOriginalKind return a slice with the original kind of array

	Array[int] will return []int
*/
func (l *Array[T]) ToOriginalKind() (res []T) {
	res = make([]T, 0)

	for i := range *l {
		res = append(res, (*l)[i])
	}

	return
}

// At return an element based on index
// accepts negative index, representing decend way
func (l *Array[T]) At(index int) (res *T) {
	// if the index passed is greater than elements count will return nil
	if index >= len(*l) {
		return
	}

	//If the index is negative when will count back forward
	if index < 0 {
		for index < 0 {
			index = len(*l) + index
		}
	}

	res = &(*l)[index]

	return
}

/*
Concat return a new Array[T] with the Arrays elements

	a := Array[int]{1, 2}
	b := Array[int]{3, 4}
	c := a.Concat(&b) // c is a new Array[int] it value is {1, 2, 3, 4}
*/
func (l *Array[T]) Concat(values ...*Array[T]) (res Array[T]) {
	res = *l

	for _, v := range values {
		for j := range *v {
			res = append(res, (*v)[j])
		}
	}

	return
}

// CopyWithin not implemented yet
func (l *Array[T]) CopyWithin() {}

// Entries not implemented yet
func (l *Array[T]) Entries() {}

// Every return true if all elements pass in the test passed by callback function
// If one element reprove the callback condition will return false
func (l *Array[T]) Every(callback func(v *T, i int) bool) bool {
	for i := range *l {
		if !callback(&(*l)[i], i) {
			return false
		}
	}

	return true
}

/*
Fill set value passed on first parameter to start (position) to the end.
If end not passed will set element to index start until the last element
*/
func (l *Array[T]) Fill(value T, start int, end ...int) *Array[T] {
	endPosition := len(*l) - 1

	if len(end) > 0 && end[0] < endPosition {
		endPosition = end[0]
	}

	for i := start; i <= endPosition; i++ {
		(*l)[i] = value
	}

	return l
}

// Filter return the elements that satisfy the callback condition
func (l *Array[T]) Filter(callback func(v *T, i int) bool) (res Array[T]) {
	for index := range *l {
		if callback(&(*l)[index], index) {
			res.Push((*l)[index])
		}
	}

	return res
}

// Find return the first element that satisfy the callback condidition
func (l *Array[T]) Find(callback func(v *T, i int) bool) (res *T) {
	index := l.FindIndex(callback)

	if index != nil {
		res = &(*l)[*index]
	}

	return
}

// FindIndex return the index of the first element that satisfy the callback condidition
// Return nil if not found any elements thath matches with the condition
func (l *Array[T]) FindIndex(callback func(v *T, i int) bool) (res *int) {
	for index := range *l {
		if callback(&(*l)[index], index) {
			res = &index
			break
		}
	}

	return
}

// FindLast return the last element that satisfy the callback condidition
func (l *Array[T]) FindLast(callback func(v *T, i int) bool) (res *T) {
	index := l.FindLastIndex(callback)

	if index != nil {
		res = &(*l)[*index]
	}

	return
}

// FindLastIndex return the index of the last element that satisfy the callback condidition
// Return nil if not found any elements thath matches with the condition
func (l *Array[T]) FindLastIndex(callback func(v *T, i int) bool) (res *int) {
	for index := len(*l) - 1; index >= 0; index-- {
		if callback(&(*l)[index], index) {
			res = &index
			break
		}
	}

	return
}

// Flat Cannot be implemented because we can't mix types inside Array
func (l *Array[T]) Flat() {}

// FlatMap Cannot be implemented because we can't mix types inside Array
func (l *Array[T]) FlatMap() {}

// ForEach loop by the Array without modificate the elements.
// But the last argument is the pointer to Array that can be modified
func (l *Array[T]) ForEach(callback func(value T, index int, array *[]T)) {
	origintalKind := l.ToOriginalKind()

	for i := range *l {
		callback((*l)[i], i, &origintalKind)
	}

	*l = AnyToArrayKind(origintalKind)
}

/*
Group return a map of the groupd elements by anything, a fields or a value
The callback function must return the value that will be used to group the elements
If the callback condition returns nil the element won't be added to any group
*/
func (l *Array[T]) Group(callback func(value T, index int) any) map[any]Array[T] {
	group := make(map[any]Array[T])

	for i := range *l {
		v := &(*l)[i]
		groupName := callback(*v, i)

		if groupName == nil {
			continue
		}

		if _, ok := group[groupName]; !ok {
			group[groupName] = Array[T]{*v}
		} else {
			group[groupName] = append(group[groupName], *v)
		}
	}

	return group
}

// Includes verify if an element exists in this Array
// If you need to verify a field of a struct, use other functions like find or filter
func (l *Array[T]) Includes(value T) bool {
	for i := range *l {
		if (*l)[i] == value {
			return true
		}
	}

	return false
}

// IndexOf return the first index of the elements that matches with the value parameter
func (l *Array[T]) IndexOf(value T) int {
	res := l.FindIndex(func(v *T, i int) bool {
		return *v == value
	})

	if res == nil {
		return -1
	}

	return *res
}

/*
Join return a string like result of joining all values by a separator

	a := Array[int]{1, 2, 3, 4, 5}
	a.Join(" - ") // will return "1 - 2 - 3 - 4 - 5"

➡ If values are structs, it needs to be the function ToString implemented, returning a string value, Example:

	type Temp struct {
		msg string
	}

	func (t *Temp) ToString() string {
		return t.msg
	}
*/
func (l *Array[T]) Join(separator string) (res string) {
	parts := make([]string, 0)

	for i := range *l {
		parts = append(parts, AnyToString(&(*l)[i]))
	}

	res = strings.Join(parts, separator)

	return
}

// Keys return all keys from a Array
func (l *Array[T]) Keys() (res []int) {
	for i := range *l {
		res = append(res, i)
	}

	return
}

// LastIndexOf return the last index of the elements that matches with the value parameter
func (l *Array[T]) LastIndexOf(value T) int {
	res := l.FindLastIndex(func(v *T, i int) bool {
		return *v == value
	})

	if res == nil {
		return -1
	}

	return *res
}

// Map iterate all elements with a callback function that can change the original value
func (l *Array[T]) Map(callback func(v *T, i int)) {
	for index := range *l {
		callback(&(*l)[index], index)
	}
}

// Pop remove the last element from this array, and return it.
// If the array is empty return nil
func (l *Array[T]) Pop() (res *T) {
	if len(*l) == 0 {
		return
	}

	// Get the index os the last element
	end := len(*l) - 1

	// Get the last element to return it
	res = &(*l)[end]

	// Remove the element from the Array
	*l = (*l)[0:end]

	return
}

// Push add onu or more elements to the end of an Array
func (l *Array[T]) Push(values ...T) {
	new := AnyToArrayKind(values)

	*l = l.Concat(&new)
}

/*
Reduce iterate all elements one by one and execute a callback that must return a value to be used on the next iteration.

➡ In the first iteration the accumulator will be nil.
If the initial values is passed, in the first iteration the accumulator value is seted with the initial value

	s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

	sumFunction := func(accumulator any, currentValue, currentIndex int) any {
		return accumulator.(int) + currentValue
	}

	result := s.Reduce(sumFunction, 0) // Result will be 0 + 1 + 2 + 3 + 4 + 5 = 15

➡ If you want right to left use ReduceRight func
*/
func (l *Array[T]) Reduce(callback func(accumulator any, currentValue T, currentIndex int) any, initialValue ...any) (accumulated any) {
	if len(initialValue) > 0 {
		accumulated = initialValue[0]
	}

	for i, v := range *l {
		accumulated = callback(accumulated, v, i)
	}

	return
}

/*
ReduceRight iterate all elements one by one right to left and execute a callback that must return a value to be used on the next iteration.

➡ In the first iteration the accumulator will be nil.
If the initial values is passed, in the first iteration the accumulator value is seted with the initial value

	s := arrayFuncs.Array[int]{1, 2, 3, 4, 5}

	sumFunction := func(accumulator any, currentValue, currentIndex int) any {
		return accumulator.(int) + currentValue
	}

	result := s.ReduceRight(sumFunction, 0) // Result will be 0 + 1 + 2 + 3 + 4 + 5 = 15

➡ If you want left to right use Reduce func
*/
func (l *Array[T]) ReduceRight(callback func(accumulator any, currentValue T, currentIndex int) any, initialValue ...any) (accumulated any) {
	if len(initialValue) > 0 {
		accumulated = initialValue[0]
	}

	// for i, v := range *l {
	for i := len(*l) - 1; i >= 0; i-- {
		accumulated = callback(accumulated, (*l)[i], i)
	}

	return
}

func (l *Array[T]) Reverse() {}

func (l *Array[T]) Shift() {}

func (l *Array[T]) Slice() {}

func (l *Array[T]) Some() {}

func (l *Array[T]) Sort() {}

func (l *Array[T]) Splice() {}

func (l *Array[T]) ToLocaleString() {}

func (l *Array[T]) ToString() {}

func (l *Array[T]) Unshift() {}

func (l *Array[T]) Values() {}
