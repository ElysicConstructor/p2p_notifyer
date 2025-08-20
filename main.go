package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	laddr, _ := net.ResolveUDPAddr("udp", ":5555")
	conn, _ := net.ListenUDP("udp", laddr)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, _ := conn.ReadFromUDP(buf)
		msg := string(buf[:n])
		if msg == "PING" {
			// Desktop-Benachrichtigung via notify-send
			exec.Command("notify-send", "P2P Anfrage", fmt.Sprintf("Peer %s möchte verbinden", addr)).Run()
		}
	}
}
