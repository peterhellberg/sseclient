package sseclient

import "testing"

func TestEvent(t *testing.T) {
	var (
		name = "foo"
		id   = "bar"
		baz  = 123
		data = map[string]interface{}{
			"baz": baz,
		}
	)

	e := Event{Name: name, ID: id, Data: data}

	if got, want := e.Name, name; got != want {
		t.Fatalf("e.Name = %q, want %q", got, want)
	}

	if got, want := e.ID, id; got != want {
		t.Fatalf("e.ID = %q, want %q", got, want)
	}

	if got, want := e.Data["baz"].(int), baz; got != want {
		t.Fatalf(`e.Data["baz"].(int) = %d, want %d`, got, want)
	}
}
