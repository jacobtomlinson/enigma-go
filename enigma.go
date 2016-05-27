package enigma

import (
  "fmt"
  "strings"
)

type Rotor struct {
  Type string
  Ring int
  Position int
}

type Settings struct {
  Rotors []Rotor
  Plugboard [26]int
  Reflector string
  Spacing int
}

func ProcessLetter(letter rune, settings *Settings, config *Config) rune {
  var letterDelta int

  if ! strings.Contains(config.ValidInput,string(letter)){
    fmt.Println("Invalid character")
  } else {

    // Iterate machine
    iterate(settings, config)

    letterDelta = runeToDelta(letter) // Convert to a unicode relative index

    // Plugboard
    letterDelta = plugboard(settings.Plugboard, letterDelta)

    // Rotor 1
    letterDelta = rotor(&settings.Rotors[0], letterDelta, true, config)

    // Rotor 2
    letterDelta = rotor(&settings.Rotors[1], letterDelta, true, config)

    // Rotor 3
    letterDelta = rotor(&settings.Rotors[2], letterDelta, true, config)

    // Reflector
    letterDelta = reflector(config.Reflectors[settings.Reflector], letterDelta, config)

    // Rotor 3
    letterDelta = rotor(&settings.Rotors[2], letterDelta, false, config)

    // Rotor 2
    letterDelta = rotor(&settings.Rotors[1], letterDelta, false, config)

    // Rotor 1
    letterDelta = rotor(&settings.Rotors[0], letterDelta, false, config)

    // Plugboard
    letterDelta = plugboard(settings.Plugboard, letterDelta)
  }

  // Letter out
  return deltaToRune(letterDelta)
}

func ProcessString(settings *Settings, input string, config *Config) string {
  returnString := ""
  for _, letter := range input {
    returnString = returnString + string(ProcessLetter(letter, settings, config))
  }
  return returnString
}

func LoadSettings(rotors []Rotor, plugboard [][2]rune, reflector string, spacing int) Settings {
  var newPlugboard [26]int
  for i := 0; i < 26; i++ {
    newPlugboard[i] = 0
  }

  for _, plug := range plugboard {
    // Calculate deltas for characters in string and update newPlugboard
    newPlugboard[runeToDelta(plug[0])] = int(plug[1]) - int(plug[0])
    newPlugboard[runeToDelta(plug[1])] = int(plug[0]) - int(plug[1])
  }

  settings := Settings{
    Rotors: rotors,
    Plugboard: newPlugboard,
    Reflector: reflector,
    Spacing: 0,
  }

  return settings
}
