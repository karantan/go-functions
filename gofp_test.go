package gofp

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	isEven := func(n int) bool {
		return n%2 == 0
	}
	t.Run("should be all even", func(t *testing.T) {
		evenNumbers := []int{2, 4, 6, 8}
		assert.True(t, All(isEven, evenNumbers))
	})
	t.Run("should not be all even", func(t *testing.T) {
		numbers := []int{2, 4, 5, 8}
		assert.False(t, All(isEven, numbers))
	})

}

func TestAny(t *testing.T) {
	isEven := func(n int) bool {
		return n%2 == 0
	}
	t.Run("should contain some even numbers", func(t *testing.T) {
		evenNumbers := []int{1, 2, 3, 4, 5}
		assert.True(t, Any(isEven, evenNumbers))
	})
	t.Run("should not any even numbers", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		assert.False(t, Any(isEven, numbers))
	})

}

func TestFilterString(t *testing.T) {
	f := func(s string) bool {
		if len(s) <= 3 {
			return true
		}
		return false
	}
	want := []string{"foo", "bar"}
	got := Filter(f, []string{"barz", "foo", "bar", "fooo"})
	assert.Equal(t, want, got)
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
	assert.Equal(t, want, got)
}

func TestSumMap(t *testing.T) {
	intData := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	gotInt := SumMap(intData)
	wantInt := int64(6)
	assert.Equal(t, wantInt, gotInt)

	floatData := map[string]float64{
		"one":   1.0,
		"two":   2.0,
		"three": 3.0,
	}

	gotFloat := SumMap(floatData)
	wantFloat := 6.0
	assert.Equal(t, wantFloat, gotFloat)
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
	assert.Equal(t, want, got)
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

func TestReduceSimple(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reducer := func(i int, j int) int {
		return i + j
	}

	got := Reduce(reducer, numbers)
	want := 45
	assert.Equal(t, want, got)
}

func TestReduceEmpty(t *testing.T) {
	numbers := []int{}
	reducer := func(i int, j int) int {
		return i + j
	}

	got := Reduce(reducer, numbers)
	want := 0
	assert.Equal(t, want, got)
}

func TestReduceStructs(t *testing.T) {
	type Item struct {
		description string
		quantity    float64
		price       float64
	}

	items := []Item{
		{"eggs", 20, 0.32},
		{"milk", 3, 1.32},
		{"oil", 2, 5},
	}

	reducer := func(price float64, i Item) float64 {
		return price + i.quantity*i.price
	}

	got := Reduce(reducer, items)
	want := 20.36
	assert.Equal(t, want, got)
}

func TestForEachReduce(t *testing.T) {
	type Person struct {
		name string
		age  int
	}

	friends := []string{"Mark", "Christopher", "Luke"}
	toPerson := func(name string) Person {
		return Person{name: name, age: len(name) * 2}
	}
	oldest := func(p1, p2 Person) Person {
		if p1.age >= p2.age {
			return p1
		}
		return p2
	}

	friendPersons := ForEach(toPerson, friends)
	want := friendPersons[1] // Christopher
	got := Reduce(oldest, friendPersons)
	assert.Equal(t, want, got)
}

func TestForEachReduceCollection(t *testing.T) {
	type Collection struct {
		number int
	}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	mapper := func(s string) Collection {
		n, err := strconv.Atoi(s)
		if err != nil {
			return Collection{0}
		}
		return Collection{n}
	}
	reducer := func(i int, c Collection) int {
		return i + c.number
	}

	got := Reduce(reducer, ForEach(mapper, numbers))
	want := 45
	assert.Equal(t, want, got)
}

func TestReduceToJson(t *testing.T) {
	type Car struct {
		brand    string
		model    string
		topSpeed int
	}

	cars := []Car{
		{"BMW", "3", 160},
		{"Opel", "Astra", 150},
		{"Audi", "A8", 220},
	}

	jsonReducer := func(s string, c Car) string {
		jsonCar := fmt.Sprintf(`{"brand": "%s", "model": "%s", "top speed": %d}`, c.brand, c.model, c.topSpeed)
		if s == "" {
			return jsonCar
		}
		return s + ",\n" + jsonCar
	}

	got := Reduce(jsonReducer, cars)
	want := `{"brand": "BMW", "model": "3", "top speed": 160},
{"brand": "Opel", "model": "Astra", "top speed": 150},
{"brand": "Audi", "model": "A8", "top speed": 220}`
	assert.Equal(t, want, got)
}
