## randomg
Generate random number based on how many digits you pass in.

Most popular current rando number algorithms end up generating the same numbers,
why the same songs come up frequently when you hit shuffle, this is an attempt to
more distribute the distribution.

### Install
`go get github.com/jujhar/randomg`

### Usage example
```go
package main

import (
	"fmt"
	"github.com/jujhar/randomg"
)

func main() {
	fmt.Println(randomg.CreateInteger(19))
}
```

### License
This project is licensed under a Simplified BSD license. Please read the [LICENSE file][license].


 [license]: https://github.com/jujhar/randomg/blob/master/LICENSE
