package updatingresource_test

import (
	"fmt"
	"time"

	"github.com/sebogh/updating-resource"
)

func Example() {

	// the update function
	f := func(x interface{}) interface{} {
		if x == nil {
			return "-"
		}
		return fmt.Sprintf("%s-", x)
	}

	// creating the new resource
	r := updatingresource.NewUpdatingResource(f, 500 * time.Millisecond)

	// query the resource 6 times
	for i := 0; i <6; i++{
		time.Sleep(250 * time.Millisecond)
		x := r.Get().(string)
		fmt.Printf("%s\n", x)
	}

	// Output:
	// -
	// --
	// --
	// ---
	// ---
	// ----
}
