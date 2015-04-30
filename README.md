# natural

natural string compare

## usage

```go
package main

import (
	"fmt"
	"github.com/mattn/natural"
)

func main() {
	a := []string{
		"３",
		"1",
		"10",
		"2",
		"20",
	}
	natural.Sort(a)
	for _, s := range a {
		fmt.Println(s) // 1 2 ３ 10 20 
	}
}
```

## License

This is rewrite of [natsort](http://sourcefrog.net/projects/natsort/). So made available under the same license as zlib.

## Author

Yasuhiro Matsumoto (a.k.a mattn)
