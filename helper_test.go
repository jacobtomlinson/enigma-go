package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestDeltaToRune(t *testing.T) {
  assert.Equal(t, 0, runeToDelta('A'), "they should be equal")
  assert.Equal(t, 25, runeToDelta('Z'), "they should be equal")
}

func TestRuneToDelta(t *testing.T) {
  assert.Equal(t, 'A', deltaToRune(0), "they should be equal")
  assert.Equal(t, 'Z', deltaToRune(25), "they should be equal")
}

func TestModDelta(t *testing.T) {
  assert.Equal(t, 0, modDelta(0), "they should be equal")
  assert.Equal(t, 0, modDelta(26), "they should be equal")
}
