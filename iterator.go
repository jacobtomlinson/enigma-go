package enigma

import (
  "strings"
)

func iterate(settings *Settings, config *Config) {
  // Shift rotor 1
  settings.Rotors[0].Position = incrementLetter(settings.Rotors[0].Position, config)
  if settings.Rotors[0].Position == config.Rotors[settings.Rotors[0].Type].Step {
    // Shift rotor 2
    settings.Rotors[1].Position = incrementLetter(settings.Rotors[1].Position, config)
    if settings.Rotors[1].Position == config.Rotors[settings.Rotors[1].Type].Step {
      // Shift rotor 3
      settings.Rotors[2].Position = incrementLetter(settings.Rotors[2].Position, config)
    }
  }
}

func incrementLetter(letter string, config *Config) string{
  index := (strings.Index(config.ValidInput, letter) + 1) % len(config.ValidInput)
  return config.ValidInput[index:index+1]
}
