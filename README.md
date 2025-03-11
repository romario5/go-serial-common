## Installing
``` bash
go get github.com/romario5/go-serial-common
```

## Usage
``` go
package main

import (
    "fmt"
    serial "github.com/romario5/go-serial-common"
)

func main() {
    packet := &serial.ChannelsPacket{}
    packet.Channels[0] = 1200
    bytes := make([]byte, serial.CHANNELS_COUNT)
    err := packet.PackChannels(bytes)
    if err != nil {
        fmt.Println("Error on packing channels", err)
    }
}