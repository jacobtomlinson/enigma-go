package enigma

func plugboard(plugboard [26]int, letterIndex int) int {
  return (letterIndex + plugboard[letterIndex] + 26) % 26
}
