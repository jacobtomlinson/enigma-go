package enigma

func runeToDelta(letter rune) int {
  return int(letter) - 65
}

func deltaToRune(delta int) rune {
  return rune(delta + 65)
}

func modDelta(delta int) int {
  return (delta + 26) % 26
}
