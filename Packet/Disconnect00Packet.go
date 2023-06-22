package Packet

// import "minecraft-proxy/Utils"

// import "minecraft-proxy/Utils"

// "fmt"
// "minecraft-proxy/Utils"
// "encoding/hex"

type Disconnect00Packet struct {
	Reason  string
}

func (p *Disconnect00Packet) Build() []byte {
	text := "{\"text\":\""+p.Reason+"\"}"
	newPacket := Packet{PacketID: 0x00, Data: []byte(text)}
	return newPacket.Build()
}

