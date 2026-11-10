package main

import (
	"fmt"
	"math/rand"
)

const (
	PENDING = iota
	ACCEPTED
	REFUSED
	UNKNOWN
)

type Status string

type Protocol string

const (
	TCP   Protocol = "TCP"
	UDP   Protocol = "UDP"
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

type PacketStatus int

const (
	Accepted PacketStatus = iota
	Rejected
	Suspicious
)

func (ps PacketStatus) String() string {
	//switch ps {
	//case Accepted:
	//	return "Accepted"
	//case Rejected:
	//	return "Rejected"
	//case Suspicious:
	//	return "Suspicious"
	//default:
	//	return "Unknown"
	//}

	// using the spread operator
	return [...]string{"Accepted", "Rejected", "Suspicious"}[ps]
}

type Packet struct {
	ID       int
	Protocol Protocol
	SrcIP    string
	DstIP    string
	Size     int
}

type PacketAnalyzer struct {
	TotalPackets      int
	AcceptedPackets   int
	RejectedPackets   int
	SuspiciousPackets int
}

func (pa *PacketAnalyzer) anaylse(p Packet) PacketStatus {
	pa.TotalPackets++

	if p.Protocol == HTTP {
		pa.SuspiciousPackets++
		return Suspicious
	}

	if p.Size > 1500 {
		pa.RejectedPackets++
		return Rejected
	}

	pa.AcceptedPackets++
	return Accepted
}

func generatePacket() Packet {
	protocols := []Protocol{TCP, UDP, HTTP, HTTPS}
	return Packet{
		ID:       rand.Intn(10000),
		Protocol: protocols[rand.Intn(len(protocols))],
		SrcIP:    fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		DstIP:    fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		Size:     rand.Intn(2000),
	}
}

func main() {
	//fmt.Println(PENDING)

	//PENDING Status = "pending"
	//ACCEPTED Status = "accepted"

	analyzer := PacketAnalyzer{}

	for i := 0; i < 100; i++ {
		packet := generatePacket()
		status := analyzer.anaylse(packet)

		fmt.Printf("Packet #%d: Protocol: %s, SrcIP: %s, DstIP: %s, Size: %d, Status: %s\n",
			packet.ID, packet.Protocol, packet.SrcIP, packet.DstIP, packet.Size, status)
	}

	fmt.Printf("\nAnalysis Summary:\n")
	fmt.Printf("Total Packets: %d\n", analyzer.TotalPackets)
	fmt.Printf("Accepted Packets: %d\n", analyzer.AcceptedPackets)
	fmt.Printf("Rejected Packets: %d\n", analyzer.RejectedPackets)
	fmt.Printf("Suspicious Packets: %d\n", analyzer.SuspiciousPackets)
}
