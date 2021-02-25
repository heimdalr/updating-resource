# updating-resource

[![Tests](https://github.com/heimdalr/updating-resource/workflows/Tests/badge.svg)](https://github.com/heimdalr/updating-resource/actions?query=workflow%3ATests)
[![Go Reference](https://pkg.go.dev/badge/github.com/heimdalr/updating-resource.svg)](https://pkg.go.dev/github.com/heimdalr/updating-resource)
[![Go Playground](https://img.shields.io/badge/go-playground-blue)](https://play.golang.org/p/l0PAgFZQ8zV)
[![Go Report Card](https://goreportcard.com/badge/github.com/heimdalr/updating-resource)](https://goreportcard.com/report/github.com/heimdalr/updating-resource)

updating-resource provides means to wrap objects whose value is then regularly and asynchronously computed / updated.

## Example

~~~~ .go
package main

import (
	"fmt"
	"time"

	"github.com/heimdalr/updating-resource"
)

func main() {

	// the function to call
	var update = func(x interface{}) (interface{}, error) {
		y := x
		if y == nil { y = "" }
		return fmt.Sprintf("%s-", y), nil
	}

	// the resource config
	config := updatingresource.Config{
		Name:     "dashes",
		Update:   update,
		Interval: 20 * time.Millisecond,
	}
	resource := config.NewResource()

	fmt.Printf("%s\n", resource.Get().(string))

	time.Sleep(30 * time.Millisecond)
	fmt.Printf("%s\n", resource.Get().(string))

	time.Sleep(20 * time.Millisecond)
	fmt.Printf("%s\n", resource.Get().(string))

}
~~~~
