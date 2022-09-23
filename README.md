## Generic List for Go.


[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)

## Contents

- [Installation](#installation)
- [Usage](#usage)
  - [From](#from)
  - [FromString](#fromstring)
  - [Range](#range)
  - [L](#l)
  - [Var](#var)
  - [Values](#values)
  - [First](#first)
  - [Last](#last)
  - [Len](#len)
  - [Add](#add)
  - [Delete](#delete)
  - [Filter](#filter)
  - [Each](#each)
  - [Chunk](#chunk)
  - [Join](#join)
  - [Nth](#nth)
  - [Random](#random)
  - [Contains](#contains)
  - [Reverse](#reverse)
  - [Shuffle](#shuffle)
  - [Unique](#unique)
  - [Zip](#zip)
  - [JoinToString](#jointostring)
  - [Fill](#fill)
  - [Sequence](#sequence)
- [Testing](#testing)
- [License](#license)


## Installation

```bash
go get -u github.com/kafkiansky/golist
```

## Usage

### From

Create the `List[V]` from given slice:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.From([]int{1, 2, 3}))
}
```

### FromString

Create the `List[string]` from given string, split it by separator and apply mapper:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.FromString[int]("1, 2, 3", ",", func(s string) (int, bool) {
        // do conversation logic here.
	}))
}
```

### Range

Create the `List[V]` from the given range:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Range(0, 10).Values()) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
}
```

### L

Alias to `From`.

### Var

Create the `List[V]` from variadic of type `V`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Values()) // [1, 2, 3]
}
```

### Values

Get the values as builtin slice `[]V` from the `List[V]`.

### First

Get the first element of `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).First()) // 1
}
```

### Last

Get the last element of `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Last()) // 3
}
```

### Len

Get the  actual len of `List[V]`.

### Add

Add the element to `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Add(4).Values()) // [1, 2, 3, 4]
}
```

### Delete

Delete the element from `List[V]` by index:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Delete(0).Values()) // [2, 3]
}
```

### Filter

Filter `List[V]` by the given filter function:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Filter(func(v int) bool {
		// do your filter logic here
		return false
    }).Values())
}
```

### Each

Iterate `List[V]` and apply a function:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Each(func(v int) int {
		// do your logic here
		return v * 2
    }).Values()) // [2, 4, 6]
}
```

Or the `Each` function that take the `List[V]` and output new `List[E]`:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/golist"
  "strconv"
)

func main() {
  fmt.Println(golist.Each(golist.Var(1, 2, 3), func(v int) string {
    return strconv.Itoa(v)
  }).Values()) // ["1", "2", "3"]
}
```

### Chunk

Chunk `List[V]` to the slice of `List[V]` by the given size:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Chunk(2)[0].Values()) // [1, 2]
}
```

### Join

Chunk `List[V]` to the slice of `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(1, 2, 3).Join(golist.Var(4, 5)).Values()) // [1, 2, 3, 4, 5]
}
```

### Nth

Get each `nth` element from the `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Nth(2).Values()) // [0, 2, 4, 6]
}
```

### Random

Get random value from the `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Random()) // 2, in example
}
```

### Contains

Check if element exists in the `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Contains(2)) // true
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Contains(11)) // false
}
```

### Reverse

Generate the reverse list of `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Reverse().Values()) // [6, 5, 4, 3, 2, 1, 0]
}
```

### Shuffle

Shuffle the given `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 1, 2, 3, 4, 5, 6).Shuffle().Values()) // [4, 1, 0, 6, 2, 5, 3], in example
}
```

### Unique

Create the `List[V]` only from unique values of the target `List[V]`:

```go
package main

import (
	"fmt"
	"github.com/kafkiansky/golist"
)

func main() {
	fmt.Println(golist.Var(0, 0, 1, 1, 2).Unique().Values()) // [0, 1, 2]
}
```

### Zip

Zip target `List[V]` with the other:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/golist"
  "log"
)

func main() {
  l, err := golist.Var(0, 1, 2).Zip(golist.Var(3, 5, 6))
  if err != nil {
    log.Fatalln(err) // occurred when lengths of the given list are different
  }

  fmt.Println(l[0].Values()) // [0, 3]
  fmt.Println(l[1].Values()) // [1, 5]
  fmt.Println(l[2].Values()) // [2, 6]
}
```

### JoinToString

Joins the `List[V]` to string if `V` is type of string.

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/golist"
)

func main() {
  fmt.Println(golist.Var("first", "second").JoinToString(", ")) // "first, second"
}
```

### Fill

Generate `List[V]` of the given `V` and provided count:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/golist"
)

func main() {
  fmt.Println(golist.Fill("?", 3).Values()) // [?, ?, ?]
}
```

### Sequence

Generate sequence. Useful for SQL query building:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/golist"
)

func main() {
  fmt.Println(golist.Sequence("$", 3, 1).JoinToString(", ")) // "$1, $2, $3"
}
```

## Testing

``` bash
$ make check
```  

## License

The MIT License (MIT). See [License File](LICENSE) for more information.