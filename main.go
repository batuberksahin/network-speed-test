package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var authors = []string{
	"batuberksahin",
}

const version = "1.0.0"

var (
	// ethernetDevice string
	packetFile string
	// output         string
)

func init() {
	// flag.StringVar(&ethernetDevice, "e", "", "Ethernet Device")
	flag.StringVar(&packetFile, "f", "", "PCAP destination")
	// flag.StringVar(&output, "o", "", "Output file destination")
}

var (
	err    error
	handle *pcap.Handle
)

func main() {
	flag.Parse()

	if packetFile == "" {
		fmt.Println("Please select a pcap!")
		return
	}

	handle, err = pcap.OpenOffline(packetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	var counter int = 0
	var length int = 0

	var start time.Time
	var current time.Time

	for packet := range packetSource.Packets() {
		if current.IsZero() {
			start = packet.Metadata().Timestamp
			current = start
		} else {
			current = packet.Metadata().Timestamp
		}

		length += packet.Metadata().Length
		counter++
	}

	elapsed := current.Sub(start)
	speed := float64(counter) / elapsed.Seconds()
	packetLength := float64(length) / float64(counter)

	fmt.Printf("PCAP took %s\n", elapsed)
	fmt.Printf("%.2f packet per second\n", speed)
	fmt.Printf("Average packet length is %.2f\n", packetLength)
}
