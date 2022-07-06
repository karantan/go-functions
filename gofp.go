package gofp

// Filter keeps elements that satisfy the test (i.e. given `f` function).
// E.g. Filter(isEven, []int{1, 2, 3, 4, 5, 6}) -> []int{2, 4, 6}
// See https://package.elm-lang.org/packages/elm/core/latest/List#filter
func Filter[T any](f func(T) bool, slice []T) []T {
	fltd := []T{}
	for _, e := range slice {
		if f(e) {
			fltd = append(fltd, e)
		}
	}
	return fltd
}

// FilterForEach filter out certain values. For example, maybe you have a bunch of
// strings from an untrusted source and you want to turn them into numbers.
// E.g. []string{"3", "hi", "12", "4th", "May"} -> []int{3, 12}
// See https://package.elm-lang.org/packages/elm/core/latest/List#filterMap
func FilterForEach[K any, T any](f func(T) (K, error), slice []T) []K {
	fltd := []K{}
	for _, e := range slice {
		v, err := f(e)
		if err == nil {
			fltd = append(fltd, v)
		}
	}
	return fltd
}

// ForEach (aka map) on slice, that will execute a function on each element of slice and
// return a new slice.
// See https://package.elm-lang.org/packages/elm/core/latest/List#map
func ForEach[T any](f func(ele T) T, slice []T) []T {
	mapped := []T{}
	for _, ele := range slice {
		mapped = append(mapped, f(ele))
	}
	return mapped
}

// Member checks if an `element` exists in the given `slice`. Returns true otherwise false.
// See https://package.elm-lang.org/packages/elm/core/latest/List#member
func Member[T comparable](element T, slice []T) bool {
	for _, v := range slice {
		if element == v {
			return true
		}
	}
	return false
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Sum list of elements.
// See https://package.elm-lang.org/packages/elm/core/latest/List#sum
func Sum[V Number](slice []V) V {
	var s V
	for _, v := range slice {
		s += v
	}
	return s
}

// SumMap sums the values of map m. It supports types that are comparable.
// See https://go.dev/ref/spec#Type_constraints and https://go.dev/ref/spec#Comparison_operators
func SumMap[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Product of the list elements.
// See https://package.elm-lang.org/packages/elm/core/latest/List#product
func Product[V Number](slice []V) V {
	if len(slice) == 0 {
		return 0
	}
	s := slice[0]
	for _, v := range slice[1:] {
		s *= v
	}
	return s
}
