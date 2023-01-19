package arrayfuncs

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

func (l *Array[T]) CopyWithin() {}

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
func (l *Array[T]) Filter(callback func(v *T, i int) bool) (res []T) {
	res = make([]T, 0)

	for index := range *l {
		if callback(&(*l)[index], index) {
			res = append(res, (*l)[index])
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

// Find return the last element that satisfy the callback condidition
func (l *Array[T]) FindLast(callback func(v *T, i int) bool) (res *T) {
	index := l.FindLastIndex(callback)

	if index != nil {
		res = &(*l)[*index]
	}

	return
}

// FindIndex return the index of the last element that satisfy the callback condidition
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

func (l *Array[T]) Join() {}

func (l *Array[T]) Keys() {}

func (l *Array[T]) LastIndexOf() {}

func (l *Array[T]) Map(callback func(i int, v *T)) {
	for index := range *l {
		callback(index, &(*l)[index])
	}
}

func (l *Array[T]) Pop() {}

func (l *Array[T]) Push() {}

func (l *Array[T]) Reduce() {}

func (l *Array[T]) ReduceRight() {}

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
