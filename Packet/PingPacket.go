package Packet



type PingPacket struct {
	time uint64
}

func IsPingPacket(packet []byte) bool {
	return packet[1] == 0x01
}
