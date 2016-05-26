package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPlugboard(t *testing.T) {
  settings := SetupTestSettings()

  assert.Equal(t, plugboard(settings.Plugboard, "A"), "B", "they should be equal")
  assert.Equal(t, plugboard(settings.Plugboard, "D"), "C", "they should be equal")
  assert.Equal(t, plugboard(settings.Plugboard, "G"), "H", "they should be equal")
  assert.Equal(t, plugboard(settings.Plugboard, "H"), "G", "they should be equal")
  assert.Equal(t, plugboard(settings.Plugboard, "Z"), "Z", "they should be equal")
}
