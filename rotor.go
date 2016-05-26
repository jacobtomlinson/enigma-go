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

func letterMaths(letter string, number int, config *Config) string{
    l_num := strings.Index(config.ValidInput, letter)
    l_num += 26
    l_num += number
    l_num %= 26
    letter = config.ValidInput[l_num:l_num+1]
    return letter
  }
