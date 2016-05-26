package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLetterMaths(t *testing.T) {
  config := getConfig()

  assert.Equal(t, "B", letterMaths("A", 1, &config), "they should be equal")
  assert.Equal(t, "A", letterMaths("B", -1, &config), "they should be equal")
  assert.Equal(t, "Z", letterMaths("A", -1, &config), "they should be equal")
  assert.Equal(t, "M", letterMaths("C", 10, &config), "they should be equal")
  assert.Equal(t, "A", letterMaths("A", 26, &config), "they should be equal")
  assert.Equal(t, "A", letterMaths("A", 104, &config), "they should be equal")
}
