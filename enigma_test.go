package enigma

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func SetupTestSettings() Settings {
  return LoadSettings(
    []Rotor{
      Rotor {
        Type: "III",
        Ring: 0,
        Position: 0, // A
      },
      Rotor {
        Type: "II",
        Ring: 0,
        Position: 0, // A
      },
      Rotor {
        Type: "I",
        Ring: 0,
        Position: 0, // A
      },
    },
    [][2]rune {
      {'A', 'B'},
      {'C', 'D'},
      {'E', 'F'},
      {'G', 'H'},
    },
    "B",
    0,
  )
}

func TestSettings(t *testing.T) {

  settings := SetupTestSettings()

  assert.Equal(t, settings.Reflector, "B", "they should be equal")
  assert.Equal(t, settings.Spacing, 0, "they should be equal")
  assert.Equal(t, settings.Plugboard[0], (int('B')-int('A')), "they should be equal")
  assert.Equal(t, settings.Rotors[0].Type, "III", "they should be equal")
  assert.Len(t, settings.Rotors, 3, "Should have three rotors")

  settings.Reflector = "A"

  assert.Equal(t, settings.Reflector, "A", "they should be equal")
}

func TestSingleLetter(t *testing.T) {

  settings := SetupTestSettings()
  config := GetConfig()

  t.Log("Testing letter H")

  the_letter := ProcessLetter('H', &settings, &config)
  assert.Equal(t, 'X', the_letter, "they should be equal")

  the_letter = ProcessLetter('E', &settings, &config)
  assert.Equal(t, 'K', the_letter, "they should be equal")
}

// func TestFlow(t *testing.T) {
//   settings := SetupTestSettings()
//   config := GetConfig()
//   letter := "H"
//
//   iterate(&settings, &config)
//
//   // Plugboard
//   letter = plugboard(settings.Plugboard, letter)
//   assert.Equal(t, "G", letter, "they should be equal")
//
//   // Rotor 1
//   letter = rotor(&settings.Rotors[0], letter, true, &config)
//   assert.Equal(t, "O", letter, "they should be equal")
//
//   // Rotor 2
//   letter = rotor(&settings.Rotors[1], letter, true, &config)
//   assert.Equal(t, "M", letter, "they should be equal")
//
//   // Rotor 3
//   letter = rotor(&settings.Rotors[2], letter, true, &config)
//   assert.Equal(t, "O", letter, "they should be equal")
//
//   // Reflector
//   letter = reflector(config.Reflectors[settings.Reflector], letter, &config)
//   assert.Equal(t, "M", letter, "they should be equal")
//
//   // Rotor 3
//   letter = rotor(&settings.Rotors[2], letter, false, &config)
//   assert.Equal(t, "C", letter, "they should be equal")
//
//   // Rotor 2
//   letter = rotor(&settings.Rotors[1], letter, false, &config)
//   assert.Equal(t, "P", letter, "they should be equal")
//
//   // Rotor 1
//   letter = rotor(&settings.Rotors[0], letter, false, &config)
//   assert.Equal(t, "X", letter, "they should be equal")
//
//   // Plugboard
//   letter = plugboard(settings.Plugboard, letter)
//   assert.Equal(t, "X", letter, "they should be equal")
// }

func TestCharsSeparately(t *testing.T) {

  settings := SetupTestSettings()
  config := GetConfig()

  t.Log("Testing HELLOWORLD individually")

  assert.Equal(t, 'X', ProcessLetter('H', &settings, &config), "they should be equal")
  assert.Equal(t, 'K', ProcessLetter('E', &settings, &config), "they should be equal")
  assert.Equal(t, 'A', ProcessLetter('L', &settings, &config), "they should be equal")
  assert.Equal(t, 'C', ProcessLetter('L', &settings, &config), "they should be equal")
  assert.Equal(t, 'B', ProcessLetter('O', &settings, &config), "they should be equal")
  assert.Equal(t, 'B', ProcessLetter('W', &settings, &config), "they should be equal")
  assert.Equal(t, 'M', ProcessLetter('O', &settings, &config), "they should be equal")
  assert.Equal(t, 'T', ProcessLetter('R', &settings, &config), "they should be equal")
  assert.Equal(t, 'B', ProcessLetter('L', &settings, &config), "they should be equal")
  assert.Equal(t, 'F', ProcessLetter('D', &settings, &config), "they should be equal")
}

func TestStrings(t *testing.T) {
  config := GetConfig()

  settings := SetupTestSettings()
  input := "HELLOWORLD"
  output := ProcessString(&settings, input, &config)
  assert.Equal(t, "XKACBBMTBF", output, "they should be equal")
  assert.Equal(t, len(input), len(output), "Should be the same length")

  settings = SetupTestSettings()
  input = output
  output = ProcessString(&settings, input, &config)
  assert.Equal(t, "HELLOWORLD", output, "they should be equal")
  assert.Equal(t, len(input), len(output), "Should be the same length")
}

func TestStringEncodings(t *testing.T) {
  config := GetConfig()
  settings := SetupTestSettings()
  assert.Equal(t, "OKPEXQVRGQJFNCFW",
               ProcessString(&settings, "TESTTESTTESTTEST", &config),
               "they should be equal")

  settings = SetupTestSettings()
  assert.Equal(t, "TESTTESTTESTTEST",
              ProcessString(&settings, "OKPEXQVRGQJFNCFW", &config),
              "they should be equal")

  settings = SetupTestSettings()
  assert.Equal(t, "BCJAHNTLJWBRXSNAXORSTNDEMFCGNUNYNTWSQYPBJDKDZFJUCSIU",
               ProcessString(&settings, "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ", &config),
               "they should be equal")

  settings = SetupTestSettings()
  assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
              ProcessString(&settings, "BCJAHNTLJWBRXSNAXORSTNDEMFCGNUNYNTWSQYPBJDKDZFJUCSIU", &config),
              "they should be equal")
}

func TestAlternateRotors(t *testing.T) {
  config := GetConfig()
  settings := SetupTestSettings()
  settings.Rotors[0].Type = "IV"
  assert.Equal(t, "HELLOWORLD",
               ProcessString(&settings, "NFGSZVMIJQ", &config),
               "they should be equal")

   settings = SetupTestSettings()
   settings.Rotors[0].Type = "V"
   settings.Rotors[1].Type = "IV"
   assert.Equal(t, "HELLOWORLD",
                ProcessString(&settings, "OZHADGADIO", &config),
                "they should be equal")
}
