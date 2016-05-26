package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestIncrementLetter(t *testing.T) {
  config := getConfig()

  assert.Equal(t, incrementLetter("A", &config), "B", "they should be equal")
  assert.Equal(t, incrementLetter("S", &config), "T", "they should be equal")
  assert.Equal(t, incrementLetter("Z", &config), "A", "they should be equal")
}

func TestIterateRotors(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()

  assert.Equal(t, settings.Rotors[0].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[1].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[2].Position, "A", "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, settings.Rotors[0].Position, "B", "they should be equal")
  assert.Equal(t, settings.Rotors[1].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[2].Position, "A", "they should be equal")

  for i := 0; i < 21; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, settings.Rotors[0].Position, "W", "they should be equal")
  assert.Equal(t, settings.Rotors[1].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[2].Position, "A", "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, settings.Rotors[0].Position, "X", "they should be equal")
  assert.Equal(t, settings.Rotors[1].Position, "B", "they should be equal")
  assert.Equal(t, settings.Rotors[2].Position, "A", "they should be equal")

  for i := 0; i < 104; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "X", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "F", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 25; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, "W", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "F", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "A", settings.Rotors[2].Position, "they should be equal")

  iterate(&settings, &config)

  assert.Equal(t, "X", settings.Rotors[0].Position, "they should be equal")
  assert.Equal(t, "G", settings.Rotors[1].Position, "they should be equal")
  assert.Equal(t, "B", settings.Rotors[2].Position, "they should be equal")

  for i := 0; i < 17423; i++ {
    iterate(&settings, &config)
  }

  assert.Equal(t, settings.Rotors[0].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[1].Position, "A", "they should be equal")
  assert.Equal(t, settings.Rotors[2].Position, "A", "they should be equal")
}
