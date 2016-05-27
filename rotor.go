package enigma

func rotor(rotor *Rotor, letter int, direction bool, config *Config) int {
  rotorMap := config.Rotors[rotor.Type].Map
  rotation := rotor.Position
  ringOffset := rotor.Ring

  if direction {
    letter = letter + rotorMap[modDelta(letter + rotation + ringOffset)][0]
  } else {
    letter = letter + rotorMap[modDelta(letter + rotation + ringOffset)][1]
  }

  return modDelta(letter)
}
