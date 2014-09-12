package sseclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	e := Event{Name: "foo", ID: "bar", Data: map[string]interface{}{
		"baz": 123,
	}}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.ID, "bar")
	assert.Equal(t, e.Data["baz"].(int), 123)
}
