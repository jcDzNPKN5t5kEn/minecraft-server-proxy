package Packet

import (
	// "fmt"
	// "encoding/hex"
)

type LoginSuccessPacket struct {
	UUID []byte
	Username string
}


func IsLoginSuccessPacket(packet []byte) bool {
	return len(packet) > 1 && packet[1] == 0x02
}
func ParseLoginSuccessPacket(packet []byte) LoginSuccessPacket {
	returnPacket := LoginSuccessPacket{
		UUID: packet[2:18],
		Username: string(packet[19:35]),
	}
	return returnPacket
}
