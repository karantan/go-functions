package gofp

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterString(t *testing.T) {
	f := func(s string) bool {
		if len(s) <= 3 {
			return true
		}
		return false
	}
	want := []string{"foo", "bar"}
	got := Filter(f, []string{"barz", "foo", "bar", "fooo"})
	assert.Equal(t, got, want)
}

func TestFilterInt(t *testing.T) {
	f := func(s int) bool {
		if s <= 3 {
			return true
		}
		return false
	}
	want := []int{1, 2, 3}
	got := Filter(f, []int{1, 2, 3, 4, 5})
	assert.Equal(t, got, want)
}

func TestSumMap(t *testing.T) {
	intData := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	gotInt := SumMap(intData)
	wantInt := int64(6)
	assert.Equal(t, gotInt, wantInt)

	floatData := map[string]float64{
		"one":   1.0,
		"two":   2.0,
		"three": 3.0,
	}

	gotFloat := SumMap(floatData)
	wantFloat := 6.0
	assert.Equal(t, gotFloat, wantFloat)
}

func TestMember(t *testing.T) {
	assert.Equal(t, Member("x", []string{"a", "b", "c"}), false)
	assert.Equal(t, Member("b", []string{"a", "b", "c"}), true)
	assert.Equal(t, Member(1, []int{1, 2, 3}), true)
	assert.Equal(t, Member(5, []int{1, 2, 3}), false)
}

func TestForEach(t *testing.T) {
	timesTwo := func(el int, i int, l []int) {
		l[i] = el * 2
	}
	testData := []int{1, 2, 3}
	ForEach(timesTwo, testData)
	assert.Equal(t, testData, []int{2, 4, 6})
}

func TestFilterForEach(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		FilterForEach(strconv.Atoi, []string{"3", "hi", "12", "4th", "May"}),
		[]int{3, 12},
	)

	f := func(i int) (string, error) {
		return strconv.Itoa(i), nil
	}

	assert.Equal(
		FilterForEach(f, []int{1, 2, 3}),
		[]string{"1", "2", "3"},
	)
}
