package enigma

import (
  "strings"
)

func letterMaths(letter string, number int, config *Config) string{
    l_num := strings.Index(config.ValidInput, letter)
    l_num += 26
    l_num += number
    l_num %= 26
    letter = config.ValidInput[l_num:l_num+1]
    return letter
}
