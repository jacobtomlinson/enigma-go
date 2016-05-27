package enigma

func reflector(reflector []int, letter int, config *Config) int {
  return letter + reflector[letter]
}
