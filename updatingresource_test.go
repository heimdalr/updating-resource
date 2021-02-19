package updatingresource_test

import (
	"fmt"
	"github.com/sebogh/updating-resource"
	"reflect"
	"testing"
	"time"
)

func TestNewResource_Basic(t *testing.T) {

	// construct an updating resource
	var updateFunction = func(x interface{}) (interface{}, error) {
		y := x
		if y == nil {
			y = ""
		}
		return fmt.Sprintf("%s-", y), nil
	}
	updateInterval := 1 * time.Second
	updatingResourceConfig := updatingresource.Config{
		Name:     "dashes",
		Update:   updateFunction,
		Interval: updateInterval,
	}
	updatingResource := updatingResourceConfig.NewResource()

	// ensure initialization worked
	if updatingResource == nil {
		t.Errorf(
			"NewResource() = nil, want %s",
			reflect.TypeOf((&updatingresource.Config{}).NewResource()).String())
	}

	// cleanup
	defer func() {
		updatingResource.Done()
	}()

	// get the value right after init
	got := updatingResource.Get()

	// type of returned object should be string
	value, ok := got.(string)
	if !ok {
		t.Errorf("expected type string got %s", reflect.TypeOf(got).String())
	}

	// we should have single dash
	if value != "-" {
		t.Errorf("expected - got %s", value)
	}

	// tick (first time)
	updatingResource.Tick()

	// give it some time to update the value
	time.Sleep(100 * time.Millisecond)

	// check (update function should have ran 2 times == f(f(nil)))
	value, _ = updatingResource.Get().(string)
	if value != "--" {
		t.Errorf("expected -- got %s", value)
	}

	// wait for the time based trigger to kick in
	time.Sleep(updateInterval)

	// check (update function should have ran 3 times == f(f(f(nil))))
	value, _ = updatingResource.Get().(string)
	if value != "---" {
		t.Errorf("expected --- got %s", value)
	}
}

func TestNewResource_SuccessCallbacks(t *testing.T) {

	// construct an updating resource
	var updateFunction = func(x interface{}) (interface{}, error) {
		y := x
		if y == nil {
			y = ""
		}
		return fmt.Sprintf("%s-", y), nil
	}
	successCounter := 0
	successFun := func(x interface{}, name string){ successCounter += 1 }
	updatingResourceConfig := updatingresource.Config{
		Name:    "dashes",
		Update:  updateFunction,
		Success: successFun,
	}
	updatingResource := updatingResourceConfig.NewResource()

	// cleanup
	defer func() {
		updatingResource.Done()
	}()

	updatingResource.Tick()

	// give it some time to update the value
	time.Sleep(100 * time.Millisecond)

	if successCounter != 2 {
		t.Errorf("expected 2; got %d", successCounter)
	}
}

func TestNewResource_ErrorCallback(t *testing.T) {

	// construct an updating resource
	var updateFunction = func(x interface{}) (interface{}, error) {
		return nil, fmt.Errorf("something bad happend")
	}
	successCounter := 0
	successFun := func(x interface{}, name string){ successCounter += 1 }
	errorCounter := 0
	errorFun := func(e error, name string){ errorCounter += 1 }
	updatingResourceConfig := updatingresource.Config{
		Name:    "dashes",
		Update:  updateFunction,
		Success: successFun,
		Error:   errorFun,
	}
	updatingResource := updatingResourceConfig.NewResource()

	// cleanup
	defer func() {
		updatingResource.Done()
	}()

	updatingResource.Tick()

	// give it some time to update the value
	time.Sleep(100 * time.Millisecond)

	if successCounter != 0 {
		t.Errorf("expected 0; got %d", successCounter)
	}
	if errorCounter != 2 {
		t.Errorf("expected 2; got %d", errorCounter)
	}

}
