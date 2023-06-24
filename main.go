package main

import (
	// "encoding/hex"
	"fmt"
	"net"

	// "encoding/binary"
	"flag"
	// "strconv"
	"minecraft-proxy/Config"
	"minecraft-proxy/Packet"
	"minecraft-proxy/Utils"
	"minecraft-proxy/WebUI"
	"minecraft-proxy/WebUI/API"
)

var (
	config        Config.Config
	remote        = flag.String("remote", "209.222.115.24:25565", "server you want to proxy")
	local         = flag.String("local", ":25565", "port to listen and forward to remote")
	webUIPort     = flag.String("webUI", ":8018", "port for WebUI")
	overwriteHost = flag.String("overwriteHost", "mc.hypixel.net", "")
	overwritePort = flag.Int("overwritePort", 25565, "")
)

func main() {
	flag.Parse()

	Config.CurConfig = Config.Config{
		Remote:        *remote,
		Local:         *local,
		WebUIPort:     *webUIPort,
		OverwriteHost: *overwriteHost,
		OverwritePort: *overwritePort,
	}

	go WebUI.Init(Config.CurConfig.WebUIPort)
	API.AddAllowedName("emotionalZombie_")
	fmt.Println("init")
	// 监听 Minecraft 客户端连接
	listener, err := net.Listen("tcp", Config.CurConfig.Local)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 处理客户端连接
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	// 读包
	buf, err := Utils.ReadConn(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if Packet.IsPingPacket(buf){
	// 	conn.Write()
	// }

	// 连接真实的 Minecraft 服务器
	serverConn, err := net.Dial("tcp", Config.CurConfig.Remote)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 替换目标服务器地址
	if Packet.IsLoggingPacket(buf) {
		Utils.PrintlnAsHex(buf, "accepted client hello: ")
		buf = modifyPacket(buf, Config.CurConfig.OverwriteHost, Config.CurConfig.OverwritePort)
		Utils.PrintlnAsHex(buf, "rebuild client hello: ")
	}

	// 发送修改后的登录包到服务器
	_, err = serverConn.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	Utils.SyncConn(conn, serverConn)

}

func modifyPacket(buf []byte, overwriteHost string, overwritePort int) []byte {
	packet := Packet.LoginPacket{
		Protocol: buf[2],
		Address:  overwriteHost,
		Port:     overwritePort,
	}
	return packet.Build()
}
