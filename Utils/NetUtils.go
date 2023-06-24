package Utils

import (
	"io"
	"net"
)

func ReadConn(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
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
			sBuf, err := ReadConn(serverConn)
			if err != nil {
				return
			}
			clientConn.Write(sBuf)
			// 将客户端的请求转发给服务器
			cBuf, err := ReadConn(serverConn)
			if err != nil {
				return
			}
			serverConn.Write(cBuf)
		}
	}()

}
