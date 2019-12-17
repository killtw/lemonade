# lemonade

A sensitive word filter service based on [DFA](https://www.wikiwand.com/en/Deterministic_finite_automaton)

## Install
```shell script
go get -u -v github.com/killtw/lemonade
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/killtw/lemonade/lemonade"
	"log"
)

func main() {
	if err := lemonade.InitTrie(); err != nil {
		log.Fatalln(err)
	}

	lemonade.Add("test")

	f1, m1 := lemonade.Replace("123test321")
	fmt.Printf("filtered: %s, matches: %s\n", f1, m1)

	f2, m2 := lemonade.Replace("123te!@#$%st321")
	fmt.Printf("filtered: %s, matches: %s\n", f2, m2)
}
```

### Output
```shell script
filtered: 123****321, matches: [test]
filtered: 123*********321, matches: [te!@#$%st]
```

## Credits

- [Karl Li](https://github.com/killtw)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
