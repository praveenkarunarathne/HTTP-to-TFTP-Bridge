package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	conn, _ := net.ListenPacket("udp", ":69")

	defer conn.Close()

	fmt.Println("TFTP server started on port 69")

	for {
		handleClient(conn)
	}
}

func handleClient(conn net.PacketConn) {
	buffer := make([]byte, 516)
	n, addr, _ := conn.ReadFrom(buffer)

	resp, err := http.Get(strings.Split(string(buffer[:n]), "\x00")[1][1:])
	if err != nil {
		fmt.Println("Error fetching HTTP URL:", err)
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.Write([]byte{0x00, 0x06})
	buf.WriteString("tsize")
	buf.WriteByte(0x00)
	buf.WriteString(strconv.Itoa(int(resp.ContentLength)))
	buf.WriteByte(0x00)

	_, _ = conn.WriteTo(buf.Bytes(), addr)

	conn.ReadFrom(buffer)

	n, addr, _ = conn.ReadFrom(buffer)

	conn.WriteTo([]byte("\x00\x06blksize\x0065464\x00"), addr)

	conn.ReadFrom(buffer)

	dataPacket := make([]byte, 4+65464)
	binary.BigEndian.PutUint16(dataPacket[:2], 3)
	blockNum := uint16(1)

	for {
		n, _ = io.ReadFull(resp.Body, dataPacket[4:])

		binary.BigEndian.PutUint16(dataPacket[2:4], blockNum)

		_, _ = conn.WriteTo(dataPacket[:n+4], addr)

		conn.ReadFrom(buffer)

		if n < 65464 {
			break
		}

		blockNum++
	}

}
