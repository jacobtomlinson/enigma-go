package enigma

func plugboard(plugboard []string, letter string) string {
  for _, plug := range plugboard {
    if letter == plug[0:1] {
      return plug[1:2]
    }
    if letter == plug[1:2] {
      return plug[0:1]
    }
  }
  return letter
}
