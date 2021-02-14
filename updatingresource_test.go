package updatingresource_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/sebogh/updating-resource"
)

func TestNewUpdatingResource(t *testing.T) {
	got := updatingresource.NewUpdatingResource("", func(x interface{}) interface{} { return 1 }, 100*time.Millisecond)
	if got == nil {
		t.Errorf("NewUpdatingResource() = nil, want %s", reflect.TypeOf(&updatingresource.UpdatingResource{}).String())
	}
}

func TestUpdatingResource_Get(t *testing.T) {
}
