package enigma

import (
  "strings"
)

func reflector(reflector string, letter string, config *Config) string {
  index := strings.Index(reflector, letter)
  return config.ValidInput[index:index+1]
}
