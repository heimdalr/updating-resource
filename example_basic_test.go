package updatingresource_test

import (
	"fmt"
	"time"

	"github.com/sebogh/updating-resource"
)

func Example() {

	// the update function
	f := func(x interface{}) interface{} {
		return fmt.Sprintf("%s-", x)
	}

	// creating the new resource
	r := updatingresource.NewUpdatingResource("-", f, 500 * time.Millisecond)

	// query the resource 8 times
	for i := 0; i <8; i++{
		time.Sleep(200 * time.Millisecond)
		x := r.Get().(string)
		fmt.Printf("%s\n", x)

		// stop updating after the 6th time
		if i == 6 {
			r.Done()
		}
	}

	// Output:
	// -
	// -
	// --
	// --
	// ---
	// ---
	// ---
	// ---
}
