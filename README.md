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
package main

import (
	"github.com/jacobtomlinson/enigma-go"
	"fmt"
)

func main() {
	config := enigma.GetConfig()
	settings := enigma.Settings{
		Rotors: []enigma.Rotor{
			enigma.Rotor{
				Type:     "I",
				Ring:     0,
				Position: 0,
			},
			enigma.Rotor{
				Type:     "II",
				Ring:     0,
				Position: 0,
			},
			enigma.Rotor{
				Type:     "III",
				Ring:     0,
				Position: 0,
			},
		},
		Plugboard: [26]int{
		},
		Reflector: "B",
		Spacing:   0,
	}

	fmt.Println(enigma.ProcessString(&settings, "HELLOWORLD", &config))

	// Outputs "MFNCZBBFZM"
}

```

## Contibuting

### Testing

```bash
go test -v github.com/jacobtomlinson/enigma-go
```
