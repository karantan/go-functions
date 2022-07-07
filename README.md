# GoFP ![gha build](https://github.com/karantan/gofp/workflows/Go/badge.svg)

A collection of helpful functions for FP (functional programming) in Go.

The goal of this package is to provide as useful and generic functions for list (slice)
manipulation.

If you want chainable functions (`.map(x -> y).reduce(x -> y, z)`) you should look at
[go_chainable](https://github.com/neurocollective/go_chainable).


## Some of supported features

### Filter
Filter keeps elements that satisfy the test.

Example:

```go
package main

import (
	"fmt"

	"github.com/karantan/gofp"
)

func main() {
	f := func(s int) bool {
		return s%2 == 0
	}
	got := gofp.Filter(f, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(got)
}
```

Output:
```go
[2, 4, 6, 8, 10]
```

### FilterForEach
FilterForEach filter out certain values. For example, maybe you have a bunch of
strings from an untrusted source and you want to turn them into numbers.

Example:

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/karantan/gofp"
)

func main() {
	got := gofp.FilterForEach(strconv.Atoi, []string{"3", "hi", "12", "4th", "May"})
	fmt.Println(got)
}
```

Output:
```go
[3, 12]
```

### ForEach
ForEach on slice will execute a function on each element of slice.

Example:

```go
package main

import (
	"fmt"

	"github.com/karantan/gofp"
)

func main() {
	timesTwo := func(el int) int{
		return el * 2
	}
	got := gofp.ForEach(timesTwo, []int{1, 2, 3})
	fmt.Println(got)
}

```

Output:
```go
[2 4 6]
```

### Member

Member checks if an `element` exists in the given `slice`. Returns true otherwise false.

Example:

```go
package main

import (
	"fmt"

	"github.com/karantan/gofp"
)

func main() {
	fmt.Println(gofp.Member("x", []string{"a", "b", "c"}))
	fmt.Println(gofp.Member("b", []string{"a", "b", "c"}))
	fmt.Println(gofp.Member(1, []int{1, 2, 3}))
	fmt.Println(gofp.Member(5, []int{1, 2, 3}))
}
```

Output:
```go
false
true
true
false
```

### SumMap

Sums the values of map m. It supports types that are comparable.

Example:

```go
package main

import (
	"fmt"

	"github.com/karantan/gofp"
)

func main() {
	intData := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(gofp.SumMap(intData))
}
```

Output:
```go
6
```

See [tests](gofp_test.go) for more examples and [go docs](https://pkg.go.dev/github.com/karantan/gofp)
for function descriptions.

## Development

To add a new feature or fix a bug, first create a new branch (or fork) and make changes
there. Then create a PR.

After merging it, you will need to create a new release tag. See [Publishing a module](https://go.dev/doc/modules/publishing)
for more details.
