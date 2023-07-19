package Utils

import (
	"fmt"
	"io"
	"net"
)

func ReadConn(conn net.Conn) ([]byte) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil
	}
	return buf[:n]
}

func SyncConn(clientConn, serverConn net.Conn) {
	// 将服务器的响应转发给客户端
	go func() {
		defer clientConn.Close()
		defer serverConn.Close()

		io.Copy(clientConn, serverConn)
	}()
	// 将客户端的请求转发给服务器
	io.Copy(serverConn, clientConn)
}

func SyncConnWrite(clientConn, serverConn net.Conn) {

	go func() {
		defer clientConn.Close()
		defer serverConn.Close()
		for {
			// 将服务器的响应转发给客户端
			sBuf := ReadConn(serverConn)
			if sBuf == nil {
				return
			}
			clientConn.Write(sBuf)
			// 将客户端的请求转发给服务器
			cBuf := ReadConn(serverConn)
			if cBuf == nil {
				return
			}
			serverConn.Write(cBuf)
		}
	}()

}
func StartProxyHTTP(conn net.Conn, targetAddr string) (bool) {
conn.Write([]byte(fmt.Sprintf("CONNECT %s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", targetAddr, targetAddr)))
		reply := ReadConn(conn)
		if reply == nil || len(string(reply)) < 14 || string(reply)[:14] != "HTTP/1.1 200 C" {
			return false
		}
		return true
}
