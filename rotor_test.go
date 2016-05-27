package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRotor(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()

  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[0], runeToDelta('A'), true, &config)), 'B', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[0], runeToDelta('B'), true, &config)), 'D', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[0], runeToDelta('Z'), true, &config)), 'O', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[1], runeToDelta('A'), true, &config)), 'A', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[1], runeToDelta('Z'), true, &config)), 'E', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[2], runeToDelta('A'), true, &config)), 'E', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[2], runeToDelta('Z'), true, &config)), 'J', "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[0], runeToDelta('A'), true, &config)), 'C', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[0], runeToDelta('D'), false, &config)), 'O', "they should be equal")
  assert.Equal(t, deltaToRune(rotor(&settings.Rotors[2], runeToDelta('O'), false, &config)), 'M', "they should be equal")
}
