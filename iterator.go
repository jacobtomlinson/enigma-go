package enigma
//
// import (
//   "strings"
// )

func iterate(settings *Settings, config *Config) {
  // Shift rotor 1
  settings.Rotors[0].Position = (settings.Rotors[0].Position + 1) % 26
  if settings.Rotors[0].Position == runeToDelta(config.Rotors[settings.Rotors[0].Type].Step) {
    // Shift rotor 2
    settings.Rotors[1].Position = (settings.Rotors[1].Position + 1) % 26
    if settings.Rotors[1].Position == runeToDelta(config.Rotors[settings.Rotors[1].Type].Step) {
      // Shift rotor 3
      settings.Rotors[2].Position = (settings.Rotors[2].Position + 1) % 26
    }
  }
}
