package updatingresource_test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sebogh/updating-resource"
)

func Example() {
	f := func(_ interface{}) interface{} { return strconv.Itoa(time.Now().Second()) }
	r := updatingresource.NewUpdatingResource(f, time.Second)

	for i := 0; i <10; i++{
		time.Sleep(300 * time.Millisecond)
		x := r.Get().(string)
		fmt.Printf("%s\n", x)
	}
}
