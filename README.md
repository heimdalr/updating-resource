# updating-resource

[![Tests](https://github.com/sebogh/updating-resource/workflows/Tests/badge.svg)](https://github.com/sebogh/updating-resource/actions?query=workflow%3ATests)
[![Go Reference](https://pkg.go.dev/badge/github.com/sebogh/updating-resource.svg)](https://pkg.go.dev/github.com/sebogh/updating-resource)

updating-resource provides means to wrap objects whose value is then regularly and asynchronously computed / updated.

## Example

~~~~ .go
package main

import (
	"fmt"
	"time"

	"github.com/sebogh/updating-resource"
)

func main() {

	// the update function
	f := func(x interface{}) interface{} {
		return fmt.Sprintf("%s-", x)
	}

	// creating the new resource
	r := updatingresource.NewUpdatingResource("-", f, 500 * time.Millisecond)

	// query the resource 6 times
	for i := 0; i <6; i++{
		time.Sleep(200 * time.Millisecond)
		x := r.Get().(string)
		fmt.Printf("%s\n", x)
	}
}
~~~~