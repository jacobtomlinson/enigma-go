package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPlugboard(t *testing.T) {
  settings := SetupTestSettings()

  assert.Equal(t, deltaToRune(plugboard(settings.Plugboard, runeToDelta('A'))), 'B', "they should be equal")
  assert.Equal(t, deltaToRune(plugboard(settings.Plugboard, runeToDelta('B'))), 'A', "they should be equal")
  assert.Equal(t, deltaToRune(plugboard(settings.Plugboard, runeToDelta('C'))), 'D', "they should be equal")
  assert.Equal(t, deltaToRune(plugboard(settings.Plugboard, runeToDelta('D'))), 'C', "they should be equal")
  assert.Equal(t, deltaToRune(plugboard(settings.Plugboard, runeToDelta('Z'))), 'Z', "they should be equal")
}
