package enigma

import (
  "strings"
)

func rotor(rotor Rotor, letter string, direction bool, config *Config) string {
  rotorMap := config.Rotors[rotor.Type].Map
  rotation := strings.Index(config.ValidInput, rotor.Position)
  ringOffset := rotor.Ring

  letter = letterMaths(letter, rotation, config)
  letter = letterMaths(letter, ringOffset, config)

  if direction {
    index := strings.Index(config.ValidInput, letter)
    letter = rotorMap[index:index+1]
  } else {
    index := strings.Index(rotorMap, letter)
    letter = config.ValidInput[index:index+1]
  }

  letter = letterMaths(letter, -rotation, config)
  letter = letterMaths(letter, -ringOffset, config)
  return letter
}
