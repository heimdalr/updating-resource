package updatingresource

import (
	"sync"
	"time"
)

// UpdatingResource is a structure to wrap an object x which is regularly and
// asynchronously computed / updated.
type UpdatingResource struct {
	mu *sync.RWMutex
	x  *interface{}
}

// NewUpdatingResource creates a new UpdatingResource. Thereby, f is the function
// that will be called every ttl to compute a new value for x (i.e. x=f(x)).
func NewUpdatingResource(f func(x interface{}) interface{}, ttl time.Duration) *UpdatingResource {
	var mu sync.RWMutex
	x := f(nil)
	go func(f func(x interface{}) interface{}) {
		ticker := time.NewTicker(ttl)
		for {
			<-ticker.C
			y := f(x)
			mu.Lock()
			x = y
			mu.Unlock()
		}
	}(f)
	resource := UpdatingResource{x: &x, mu: &mu}
	return &resource
}

// Get returns the current value of the encapsulated object. Get is thread-safe
// wrt. to the function updating the encapsulated object.
func (r *UpdatingResource) Get() interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return *r.x
}
