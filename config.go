package enigma

type Config struct {
  ValidInput string
	Reflectors map[string]string
	Rotors map[string]RotorMap
}

type RotorMap struct {
	Map string
	Step string
}

func getConfig() Config{
	rotors := make(map[string]RotorMap)
	rotors["I"] = RotorMap{ Map: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", Step: "R", }
	rotors["II"] = RotorMap{ Map: "AJDKSIRUXBLHWTMCQGZNPYFVOE", Step: "F",	}
	rotors["III"] = RotorMap{	Map: "BDFHJLCPRTXVZNYEIWGAKMUSQO", Step: "W",	}

	reflectors := make(map[string]string)
  reflectors["B"] = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
  reflectors["C"] = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
  reflectors["B Dünn"] = "ENKQAUYWJICOPBLMDXZVFTHRGS"
  reflectors["C Dünn"] = "RDOBJNTKVEHMLFCWZAXGYIPSUQ"

	return Config{
		ValidInput: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Rotors: rotors,
		Reflectors: reflectors,
	}
}
