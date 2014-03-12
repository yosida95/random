# random

## About
This is a package to generate random string.

## LICENSE
This package is licensed under the [MIT LICENS](http://yosida95.mit-license.org/).

## How to use
```go
package main

import (
	"github.com/yosida95/random"
	"log"
)

func generateRandomString(length int) (str string, err error) {
	r, err := random.Ascii(random.LOWER | random.UPPER)
	if err != nil {
		return
	}

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		r, _, _ := r.ReadRune()
		runes[i] = r
	}

	str = string(runes)
	return
}

func main() {
	password, err := generateRandomString(32)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", password)
}
```
