package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func SetupTestSettings() Settings {
  return Settings {
    Rotors: [3]Rotor{
      Rotor {
        Type: "III",
        Ring: 0,
        Position: "A",
      },
      Rotor {
        Type: "II",
        Ring: 0,
        Position: "A",
      },
      Rotor {
        Type: "I",
        Ring: 0,
        Position: "A",
      },
    },
    Plugboard: []string {
      "AB",
      "CD",
      "EF",
      "GH",
    },
    Reflector: "B",
    Spacing: 0,
  }
}

func TestSingleLetter(t *testing.T) {

  settings := SetupTestSettings()

  t.Log("Testing letter H")

  the_letter := ProcessLetter(&settings, "H")
  assert.Equal(t, "X", the_letter, "they should be equal")
  assert.Len(t, the_letter, 1, "Should return a single letter")

  the_letter = ProcessLetter(&settings, "E")
  assert.Equal(t, "K", the_letter, "they should be equal")
  assert.Len(t, the_letter, 1, "Should return a single letter")
}

func TestCharsSeparately(t *testing.T) {

  settings := SetupTestSettings()

  t.Log("Testing HELLOWORLD individually")

  assert.Equal(t, "X", ProcessLetter(&settings, "H"), "they should be equal")
  assert.Equal(t, "K", ProcessLetter(&settings, "E"), "they should be equal")
  assert.Equal(t, "A", ProcessLetter(&settings, "L"), "they should be equal")
  assert.Equal(t, "C", ProcessLetter(&settings, "L"), "they should be equal")
  assert.Equal(t, "B", ProcessLetter(&settings, "O"), "they should be equal")
  assert.Equal(t, "B", ProcessLetter(&settings, "W"), "they should be equal")
  assert.Equal(t, "M", ProcessLetter(&settings, "O"), "they should be equal")
  assert.Equal(t, "T", ProcessLetter(&settings, "R"), "they should be equal")
  assert.Equal(t, "B", ProcessLetter(&settings, "L"), "they should be equal")
  assert.Equal(t, "F", ProcessLetter(&settings, "D"), "they should be equal")
}

func TestStrings(t *testing.T) {

  settings := SetupTestSettings()

  input := "HELLOWORLD"
  t.Log("Encrypting '" + input + "'")
  output := ProcessString(&settings, input)
  t.Log("Encrypted as '" + output + "'")

  assert.Equal(t, "XKACBBMTBF", output, "they should be equal")
  assert.Equal(t, len(input), len(output), "Should be the same length")

  settings = SetupTestSettings()

  input = output
  t.Log("Decrypting '" + input + "'")
  output = ProcessString(&settings, input)
  t.Log("Decrypted as '" + output + "'")

  assert.Equal(t, "HELLOWORLD", output, "they should be equal")
  assert.Equal(t, len(input), len(output), "Should be the same length")
}

func TestSettings(t *testing.T) {

  settings := SetupTestSettings()

  t.Log("Testing settings")

  assert.Equal(t, settings.Reflector, "B", "they should be equal")
  assert.Equal(t, settings.Spacing, 0, "they should be equal")
  assert.Equal(t, settings.Plugboard[0], "AB", "they should be equal")
  assert.Equal(t, settings.Rotors[0].Type, "III", "they should be equal")
  assert.Len(t, settings.Rotors, 3, "Should have three rotors")

  settings.Reflector = "A"

  assert.Equal(t, settings.Reflector, "A", "they should be equal")
}

func TestFlow(t *testing.T) {
  settings := SetupTestSettings()
  config := getConfig()
  letter := "H"

  iterate(&settings, &config)

  // Plugboard
  letter = plugboard(settings.Plugboard, letter)
  assert.Equal(t, "G", letter, "they should be equal")

  // Rotor 1
  letter = rotor(settings.Rotors[0], letter, true, &config)
  assert.Equal(t, "O", letter, "they should be equal")

  // Rotor 2
  letter = rotor(settings.Rotors[1], letter, true, &config)
  assert.Equal(t, "M", letter, "they should be equal")

  // Rotor 3
  letter = rotor(settings.Rotors[2], letter, true, &config)
  assert.Equal(t, "O", letter, "they should be equal")

  // Reflector
  letter = reflector(config.Reflectors[settings.Reflector], letter, &config)
  assert.Equal(t, "M", letter, "they should be equal")

  // Rotor 3
  letter = rotor(settings.Rotors[2], letter, false, &config)
  assert.Equal(t, "C", letter, "they should be equal")

  // Rotor 2
  letter = rotor(settings.Rotors[1], letter, false, &config)
  assert.Equal(t, "P", letter, "they should be equal")

  // Rotor 1
  letter = rotor(settings.Rotors[0], letter, false, &config)
  assert.Equal(t, "X", letter, "they should be equal")

  // Plugboard
  letter = plugboard(settings.Plugboard, letter)
  assert.Equal(t, "X", letter, "they should be equal")
}
