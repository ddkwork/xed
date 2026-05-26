# XED Go Bindings

Go bindings for Intel XED (X86 Encoder Decoder) library.

## Installation

```bash
go get github.com/ddkwork/xed
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/ddkwork/xed"
)

func main() {
    xed := xed.NewXed()
    defer xed.GlobalShutdown()
    
    xed.GlobalInit()
    
    // Decode instruction
    var xdi xed.XedDecodedInst
    xed.Decode(&xdi, []byte{0xC3}, 1, xed.XED_MACHINE_MODE_LEGACY_32)
    
    fmt.Println(xed.Disassemble(&xdi))
}
```

## Features

- Full XED API coverage
- Instruction decoding
- Instruction encoding
- Disassembly output
- Register and operand access

## Requirements

- Windows x64
- Go 1.21+

## License

MIT License
