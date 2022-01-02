package testingv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolutev3(t *testing.T) {

	//buat subtest
	t.Run("negative case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})

	t.Run("Positive case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})

}

func TestLuasSegitigav3(t *testing.T) {
	t.Run("negative case", func(t *testing.T) {
		hitung1 := LuasSegitiga(10, 3)
		assert.Equal(t, 30, hitung1)
	})

	t.Run("Positive case", func(t *testing.T) {
		hitung2 := LuasSegitiga(10, 2)

		assert.Equal(t, 20, hitung2)
	})
}

func TestAdd(t *testing.T) {
	t.Run("negative case", func(t *testing.T) {
		n := add(-1, -2)
		assert.Equal(t, -3, n)
	})

	t.Run("positive case", func(t *testing.T) {
		n := add(1, -2)
		assert.Equal(t, -1, n)
	})
}
