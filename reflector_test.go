package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReflector(t *testing.T) {
  config := getConfig()

  assert.Equal(t, deltaToRune(reflector(config.Reflectors["B"], runeToDelta('A'), &config)), 'Y', "they should be equal")
  assert.Equal(t, deltaToRune(reflector(config.Reflectors["B"], runeToDelta('Y'), &config)), 'A', "they should be equal")

  assert.Equal(t, deltaToRune(reflector(config.Reflectors["C"], runeToDelta('I'), &config)), 'E', "they should be equal")
  assert.Equal(t, deltaToRune(reflector(config.Reflectors["C"], runeToDelta('E'), &config)), 'I', "they should be equal")

  assert.Equal(t, deltaToRune(reflector(config.Reflectors["B Dünn"], runeToDelta('Z'), &config)), 'S', "they should be equal")
  assert.Equal(t, deltaToRune(reflector(config.Reflectors["B Dünn"], runeToDelta('S'), &config)), 'Z', "they should be equal")

  assert.Equal(t, deltaToRune(reflector(config.Reflectors["C Dünn"], runeToDelta('K'), &config)), 'H', "they should be equal")
  assert.Equal(t, deltaToRune(reflector(config.Reflectors["C Dünn"], runeToDelta('H'), &config)), 'K', "they should be equal")
}
