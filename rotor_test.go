package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRotor(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()

  assert.Equal(t, rotor(settings.Rotors[0], "A", true, &config), "B", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[0], "B", true, &config), "D", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[0], "Z", true, &config), "O", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[1], "A", true, &config), "A", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[1], "Z", true, &config), "E", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[2], "A", true, &config), "E", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[2], "Z", true, &config), "J", "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, rotor(settings.Rotors[0], "A", true, &config), "C", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[0], "D", false, &config), "O", "they should be equal")
  assert.Equal(t, rotor(settings.Rotors[2], "O", false, &config), "M", "they should be equal")
}
