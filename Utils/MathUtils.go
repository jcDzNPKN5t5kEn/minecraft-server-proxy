package Utils

func PortToBytes(port int) []byte {
	portByte := []byte{byte(port >> 8), byte(port)}
	return portByte
}

