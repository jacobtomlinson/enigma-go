package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReflector(t *testing.T) {
  config := getConfig()

  assert.Equal(t, reflector(config.Reflectors["B"], "A", &config), "Y", "they should be equal")
  assert.Equal(t, reflector(config.Reflectors["B"], "Y", &config), "A", "they should be equal")

  assert.Equal(t, reflector(config.Reflectors["C"], "I", &config), "E", "they should be equal")
  assert.Equal(t, reflector(config.Reflectors["C"], "E", &config), "I", "they should be equal")

  assert.Equal(t, reflector(config.Reflectors["B Dünn"], "Z", &config), "S", "they should be equal")
  assert.Equal(t, reflector(config.Reflectors["B Dünn"], "S", &config), "Z", "they should be equal")

  assert.Equal(t, reflector(config.Reflectors["C Dünn"], "K", &config), "H", "they should be equal")
  assert.Equal(t, reflector(config.Reflectors["C Dünn"], "H", &config), "K", "they should be equal")
}
