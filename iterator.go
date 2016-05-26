package enigma

import (
  "strings"
)

func iterate(settings *Settings, config *Config) {

  shiftRotorTwo := false
  shiftRotorThree := false

  if settings.Rotors[0].Position == config.Rotors[settings.Rotors[0].Type].Step {
    shiftRotorTwo = true

    if settings.Rotors[1].Position == config.Rotors[settings.Rotors[1].Type].Step {
      shiftRotorThree = true
    }
  }

  // Shift rotor 1
  settings.Rotors[0].Position = incrementLetter(settings.Rotors[0].Position, config)

  // Shift rotor 2
  if shiftRotorTwo  {
    settings.Rotors[1].Position = incrementLetter(settings.Rotors[1].Position, config)
  }

  // Shift rotor 3
  if shiftRotorThree {
    settings.Rotors[2].Position = incrementLetter(settings.Rotors[2].Position, config)
  }
}

func incrementLetter(letter string, config *Config) string{
  index := (strings.Index(config.ValidInput, letter) + 1) % len(config.ValidInput)
  return config.ValidInput[index:index+1]
}
