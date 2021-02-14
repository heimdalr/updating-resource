# updating-resource

[![Tests](https://github.com/sebogh/updating-resource/workflows/Tests/badge.svg)](https://github.com/sebogh/updating-resource/actions?query=workflow%3ATests)
[![Go Reference](https://pkg.go.dev/badge/github.com/sebogh/updating-resource.svg)](https://pkg.go.dev/github.com/sebogh/updating-resource)

updating-resource provides means to wrap objects whose value is then regularly and asynchronously computed / updated.

## Example

~~~~ .go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sebogh/updating-resource"
)

func main() {
	f := func(_ interface{}) interface{} { return strconv.Itoa(time.Now().Second()) }
	r := updatingresource.NewUpdatingResource(f, time.Second)

	for i := 0; i <10; i++{
		time.Sleep(300 * time.Millisecond)
		x := r.Get().(string)
		fmt.Printf("%s\n", x)
	}
}
~~~~