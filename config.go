package enigma

type Config struct {
  ValidInput string
	Reflectors map[string][]int
	Rotors map[string]RotorMap
}

type RotorMap struct {
	Map [][2]int
	Step rune
}

func GetConfig() Config{
	rotors := make(map[string]RotorMap)
  alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rotors["I"] = RotorMap{ Map: rotorMapStringToArray("EKMFLGDQVZNTOWYHXUSPAIBRCJ"), Step: 'R', }
	rotors["II"] = RotorMap{ Map: rotorMapStringToArray("AJDKSIRUXBLHWTMCQGZNPYFVOE"), Step: 'F',	}
	rotors["III"] = RotorMap{	Map: rotorMapStringToArray("BDFHJLCPRTXVZNYEIWGAKMUSQO"), Step: 'W',	}
	rotors["IV"] = RotorMap{	Map: rotorMapStringToArray("ESOVPZJAYQUIRHXLNFTGKDCMWB"), Step: 'K',	}
	rotors["V"] = RotorMap{	Map: rotorMapStringToArray("VZBRGITYUPSDNHLXAWMJQOFECK"), Step: 'A',	}

	reflectors := make(map[string][]int)
  reflectors["B"] = reflectorMapStringToArray(alphabet, "YRUHQSLDPXNGOKMIEBFZCWVJAT")
  reflectors["C"] = reflectorMapStringToArray(alphabet, "FVPJIAOYEDRZXWGCTKUQSBNMHL")
  reflectors["B Dünn"] = reflectorMapStringToArray(alphabet, "ENKQAUYWJICOPBLMDXZVFTHRGS")
  reflectors["C Dünn"] = reflectorMapStringToArray(alphabet, "RDOBJNTKVEHMLFCWZAXGYIPSUQ")

	return Config{
		ValidInput: alphabet,
		Rotors: rotors,
		Reflectors: reflectors,
	}
}

func rotorMapStringToArray(mapping string) [][2]int {
  var newMapping [][2]int
  for i, mapp := range mapping {
    newMapping = append(newMapping, [2]int{runeToDelta(mapp) - i, 0})
  }
  for i, mapp := range mapping {
    newMapping[runeToDelta(mapp)][1] = i - runeToDelta(mapp)
  }
  return newMapping
}

func reflectorMapStringToArray(alphabet string, mapping string) []int {
  var newMapping []int
  for i, mapp := range mapping {
      newMapping = append(newMapping, runeToDelta(mapp) - i)
  }
  return newMapping
}
