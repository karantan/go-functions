package gofp

import (
	"reflect"
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
	timesTwo := func(el int) int {
		return el * 2
	}
	got := ForEach(timesTwo, []int{1, 2, 3})
	want := []int{2, 4, 6}
	assert.Equal(t, got, want)
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

func TestSum(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty slice", args{slice: []int{}}, 0},
		{"sum ints", args{slice: []int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	type args struct {
		slice []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty slice", args{slice: []float64{}}, 0},
		{"product ints", args{slice: []float64{1, 2, 3, 4}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Product(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Product() = %v, want %v", got, tt.want)
			}
		})
	}
}
