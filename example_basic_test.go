package updatingresource_test

import (
	"fmt"
	"time"

	"github.com/sebogh/updating-resource"
)

func Example() {

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
		Interval: 300 * time.Millisecond,
	}
	resource := config.NewResource()

	fmt.Printf("%s\n", resource.Get().(string))

	resource.Tick()
	time.Sleep(50 * time.Millisecond)

	fmt.Printf("%s\n", resource.Get().(string))

	time.Sleep(450 * time.Millisecond)

	fmt.Printf("%s\n", resource.Get().(string))

	// Output:
	// -
	// --
	// ---
}
