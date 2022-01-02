package testingv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolutev2(t *testing.T) {
	c := Absolute(-5)

	assert.Equal(t, 5, c)

}

func TestLuasSegitigav2(t *testing.T) {
	hitung := LuasSegitiga(20, 2)
	assert.Equal(t, 40, hitung)
}
