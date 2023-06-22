package Packet

import (
	// "fmt"
	"minecraft-proxy/Utils"
	// "encoding/hex"
)

type LoginPacket struct {
	Protocol byte
	Address  string
	Port     int
}

func (p *LoginPacket) Build() []byte {
	ServerAddrLen := byte(len(p.Address))
	ServerAddr := []byte(p.Address)
	portBytes := Utils.PortToBytes(p.Port)
	state := byte(0x02)
	buf := []byte{p.Protocol, ServerAddrLen}
	buf = append(buf, ServerAddr...)
	buf = append(buf, portBytes...)
	buf = append(buf, state)
	newPacket := Packet{PacketID: 0x00, Data: buf}
	newPacketBytes := newPacket.Build()
	return newPacketBytes
}

func IsLoggingPacket(packet []byte) bool {
	return packet[1] == 0x00 && packet[len(packet)-1] == 0x02
}
