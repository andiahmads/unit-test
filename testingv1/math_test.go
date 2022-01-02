package testingv1

import "testing"

func TestAbsolute(t *testing.T) {
	c := Absolute(-5)
	if c != 5 {
		t.Logf("Expect 5, but got %d", c)
		t.Fail()
	}

}

func TestLuasSegitiga(t *testing.T) {
	v := LuasSegitiga(10, 2)
	total := v
	if total != 20 {
		t.Logf("expect 20,but got %d", v)
		t.Fail()
	}
}
