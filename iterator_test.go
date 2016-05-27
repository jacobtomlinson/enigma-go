package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestIterateRotors(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()

  assert.Equal(t, runeToDelta('A'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, runeToDelta('B'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 21; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, runeToDelta('W'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('B'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, runeToDelta('X'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('B'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 104; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, runeToDelta('X'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('F'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('B'), settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 25; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, runeToDelta('W'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('G'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('B'), settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, runeToDelta('X'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('G'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('B'), settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 17423; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, runeToDelta('A'), settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, runeToDelta('A'), settings.Rotors[2].Position, "they should be equal")
}
