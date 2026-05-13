package builder

import "testing"

func TestBuild(t *testing.T) {
	t.Run("build", func(t *testing.T) {
		director := &Director{}
		v_builder := &VillaBuilder{}
		director.builder = v_builder
		villa := director.Construct()
		t.Log(villa)
	})
}
