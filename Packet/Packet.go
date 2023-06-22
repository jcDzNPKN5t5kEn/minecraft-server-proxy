package Packet

import (
	"bytes"
	"encoding/binary"
)

type Packet struct {
    Length   int
    PacketID byte
    Data     []byte
}

func (p *Packet) Build() []byte {
    // calculate the length of packet data + length of the packet ID
    dataLen := len(p.Data) + 1

    // create a buffer with initial length of dataLen+1
    buf := bytes.NewBuffer(make([]byte, 0, dataLen+1))

    // write the length to the buffer as a varint
    binary.Write(buf, binary.BigEndian, uint32(dataLen))

    // write the packet ID to the buffer
    binary.Write(buf, binary.BigEndian, p.PacketID)

    // write the data to the buffer
    buf.Write(p.Data)

    return buf.Bytes()
}

