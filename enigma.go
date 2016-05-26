package enigma

import (
  "fmt"
  "strings"
)

type Rotor struct {
  Type string
  Ring int
  Position string
}

type Settings struct {
  Rotors [3]Rotor
  Plugboard []string
  Reflector string
  Spacing int
}

func ProcessLetter(letter string, settings *Settings, config *Config) string {

  if ! strings.Contains(config.ValidInput,letter){
    fmt.Println("Invalid character")
  } else {

    // Iterate machine
    iterate(settings, config)

    // Plugboard
    letter = plugboard(settings.Plugboard, letter)

    // Rotor 1
    letter = rotor(&settings.Rotors[0], letter, true, config)

    // Rotor 2
    letter = rotor(&settings.Rotors[1], letter, true, config)

    // Rotor 3
    letter = rotor(&settings.Rotors[2], letter, true, config)

    // Reflector
    letter = reflector(config.Reflectors[settings.Reflector], letter, config)

    // Rotor 3
    letter = rotor(&settings.Rotors[2], letter, false, config)

    // Rotor 2
    letter = rotor(&settings.Rotors[1], letter, false, config)

    // Rotor 1
    letter = rotor(&settings.Rotors[0], letter, false, config)

    // Plugboard
    letter = plugboard(settings.Plugboard, letter)
  }

  // Letter out
  return letter
}

func ProcessString(settings *Settings, input string) string {
  config := getConfig()
  returnString := ""
  for i := 0; i < len(input); i++ {
    returnString = returnString + ProcessLetter(input[i:i+1], settings, &config)
  }
  return returnString
}
