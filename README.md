# Enigma [![Build Status](https://travis-ci.org/jacobtomlinson/enigma-go.svg)](https://travis-ci.org/jacobtomlinson/enigma-go)

An enigma machine emulator.

Build using the specification at http://www.codesandciphers.org.uk/enigma/index.htm.

Currently implements an M3 enigma machine using rotors I, II, III, IV and V.

## Install

```bash
go get github.com/jacobtomlinson/enigma-go
```

## Usage

```Go
import (
  "github.com/jacobtomlinson/enigma-go"
  "fmt"
)

func main(){
  settings := enigma.Settings {
    Rotors: [3]enigma.Rotor{
      enigma.Rotor {
        Type: "III",
        Ring: 0,
        Position: "A",
      },
      enigma.Rotor {
        Type: "II",
        Ring: 0,
        Position: "A",
      },
      enigma.Rotor {
        Type: "I",
        Ring: 0,
        Position: "A",
      },
    },
    Plugboard: []string {
    },
    Reflector: "C",
    Spacing: 0,
  }

  fmt.Println(enigma.ProcessString(&settings, "HELLOWORLD"))

  // Outputs "XKACBBMTBF"
}

```

## Contibuting

### Testing

```bash
go test -v github.com/jacobtomlinson/enigma-go
```
