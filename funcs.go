package arrayfuncs

type Array[T any] []T

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

func (l *Array[T]) Every() {}

func (l *Array[T]) Fill() {}

func (l *Array[T]) Filter(callback func(v *T, i int) bool) (res []T) {
	res = make([]T, 0)

	for index := range *l {
		if callback(&(*l)[index], index) {
			res = append(res, (*l)[index])
		}
	}

	return res
}

func (l *Array[T]) Find(callback func(v *T, i int) bool) (res T) {
	for index := range *l {
		if callback(&(*l)[index], index) {
			res = (*l)[index]
			break
		}
	}

	return res
}

func (l *Array[T]) FindIndex() {}

func (l *Array[T]) FindLastIndex() {}

func (l *Array[T]) Flat() {}

func (l *Array[T]) FlatMap() {}

func (l *Array[T]) ForEach() {}

func (l *Array[T]) GroupToMap() {}

func (l *Array[T]) Includes() {}

func (l *Array[T]) IndexOf() {}

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
