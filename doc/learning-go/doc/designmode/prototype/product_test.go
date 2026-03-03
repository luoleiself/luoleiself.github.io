package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	t.Run("prototype", func(t *testing.T) {
		original := &Product{name: "Phone", category: "Electronics"}
		t.Log(original.GetDetails())

		cloned := original.Clone().(*Product)
		cloned.SetName("SmartPhone")
		t.Log(cloned.GetDetails())
	})
}
