package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
  config := getConfig()

  assert.Equal(t, config.ValidInput, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "they should be equal")
  assert.Equal(t, config.Rotors["I"].Step, 'R', "they should be equal")
}
