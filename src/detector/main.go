package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

var (
	device      string = "eth1"
	snapshot    int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func getMyIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	handle, err := pcap.OpenLive(device, snapshot, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	err = handle.SetDirection(pcap.DirectionIn)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("My IP Address: " + getMyIPAddress())

	beforeTime := time.Now()
	for packet := range packetSource.Packets() {
		// パケットを受信するごとに時刻を取ってしきい値以下ならほげほげ
		currentTime := time.Now()
		sub := beforeTime.Sub(currentTime)
		beforeTime = currentTime

		fmt.Println("====================================================")
		fmt.Println(sub)
		fmt.Println(packet)
	}
}
