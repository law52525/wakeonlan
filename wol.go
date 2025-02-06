package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
  "time"
)

func wakeOnLan(mac, ip string, port int) error {
	// Clean the MAC address (remove any ':' symbols)
	mac = strings.ReplaceAll(mac, ":", "")

	// Create Magic Packet (FF * 6 + MAC address * 16)
	var magicPacket []byte
	// Add 6 bytes of 0xFF
	for i := 0; i < 6; i++ {
		magicPacket = append(magicPacket, 0xFF)
	}

	// Add the MAC address 16 times
	macBytes, err := hex.DecodeString(mac)
	if err != nil {
		return fmt.Errorf("invalid MAC address: %v", err)
	}
	for i := 0; i < 16; i++ {
		magicPacket = append(magicPacket, macBytes...)
	}

	// Resolve the broadcast address
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to send the magic packet: %v", err)
	}
	defer conn.Close()

	// Set the socket to broadcast
	conn.SetDeadline(time.Now().Add(time.Second))

	// Send the magic packet
	_, err = conn.Write(magicPacket)
	if err != nil {
		return fmt.Errorf("failed to write magic packet: %v", err)
	}

	fmt.Println("Magic Packet sent!")
	return nil
}

func main() {
	// Usage example
	macAddress := "11:22:33:44:55:66"
	ipAddress := "255.255.255.255" // Broadcaster IP address (use the broadcast IP for your subnet)
	port := 9 // Default Wake-on-LAN port

	err := wakeOnLan(macAddress, ipAddress, port)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Wake-on-LAN packet sent successfully!")
	}
	time.Sleep(3 * time.Second)
}
