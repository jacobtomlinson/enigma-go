package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestIncrementLetter(t *testing.T) {
  config := getConfig()

  assert.Equal(t, "B", incrementLetter("A", &config), "they should be equal")
  assert.Equal(t, "T", incrementLetter("S", &config), "they should be equal")
  assert.Equal(t, "A", incrementLetter("Z", &config), "they should be equal")
}

func TestIterateRotors(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()

  assert.Equal(t, "A", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, "B", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 21; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "W", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, "X", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 104; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "X", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "F", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 25; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "W", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "G", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, "X", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "G", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 17423; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "A", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")
}
