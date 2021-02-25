package updatingresource

import (
	"sync"
	"time"
)

// Config is the configuration for a Resource.
type Config struct {

	// Name of this resource.
	Name string

	// Update is the function to update the wrapped object. Update must handle nil
	// values.
	Update func(x interface{}) (interface{}, error)

	// Interval is the time to wait between between automatic updates. If 0 (or not
	// set), updates will never trigger automatically.
	Interval time.Duration

	// Success is the function (if any) to call, if the object was successfully
	// updated. Success is called with the new value the name of the resource.
	Success func(x interface{}, name string)

	// Success is the function (if any) to call, if the object was successfully
	// updated. Error is called with the the error that occurred the name of the
	// resource.
	Error func(e error, name string)
}

// Resource regularly, asynchronously updates a wrapped object.
type Resource struct {

	// resource configuration
	*Config

	// mutex for protecting object updates
	mu *sync.RWMutex

	// the wrapped object
	x *interface{}

	// a channel for terminating the updates
	done chan bool

	// channel to manually trigger an update
	tick chan bool
}

// NewResource creates a new Resource.
func (c *Config) NewResource() *Resource {
	var mu sync.RWMutex

	done := make(chan bool)
	tick := make(chan bool)

	var startValue interface{} = nil
	x := &startValue

	var timeTicker <-chan time.Time
	if c.Interval > 0 {
		timeTicker = time.NewTicker(c.Interval).C
	}

	go func(f func(x interface{}) (interface{}, error)) {

		// wrap f to be used for inputs on multiple channels
		var fWrapper = func() {
			y, err := f(*x)
			if err != nil && c.Error != nil {
				c.Error(err, c.Name)
			} else {
				if c.Success != nil {
					c.Success(y, c.Name)
				}
				mu.Lock()
				*x = y
				mu.Unlock()
			}
		}

		// wait (forever) on channels
		for {
			select {
			case <-done:
				return
			case <-timeTicker:
				fWrapper()
			case <-tick:
				fWrapper()
			}
		}
	}(c.Update)

	// compute step 0
	tick <- true

	resource := Resource{
		Config: c,
		mu:     &mu,
		x:      x,
		done:   done,
		tick:   tick,
	}
	return &resource
}

// Get returns the current value of the wrapped object. Get is thread-safe
// wrt. to the function updating the encapsulated object.
func (r *Resource) Get() interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return *r.x
}

// Done stops the updating of the wrapped object.
func (r *Resource) Done() {
	r.done <- true
}

// Tick manually trigger an update.
func (r *Resource) Tick() {
	r.tick <- true
}
