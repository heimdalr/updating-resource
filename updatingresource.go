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
	done chan bool
}

// NewUpdatingResource creates a new UpdatingResource. Thereby, f is the function
// that will be called every ttl to compute a new value for x (i.e. x=f(x)).
func NewUpdatingResource(x interface{}, f func(x interface{}) interface{}, ttl time.Duration) *UpdatingResource {
	var mu sync.RWMutex
	done := make(chan bool)
	go func(f func(x interface{}) interface{}) {
		ticker := time.NewTicker(ttl)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				y := f(x)
				mu.Lock()
				x = y
				mu.Unlock()
			}
		}
	}(f)
	resource := UpdatingResource{x: &x, mu: &mu, done: done}
	return &resource
}

// Get returns the current value of the wrapped object. Get is thread-safe
// wrt. to the function updating the encapsulated object.
func (r *UpdatingResource) Get() interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return *r.x
}

// Done stops the updating of the wrapped object.
func (r *UpdatingResource) Done() {
	r.done <- true
}