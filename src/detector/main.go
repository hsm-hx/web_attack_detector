package main

import (
  "fmt"
  "log"
  "time"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
  device string = "eth1"
  snapshot_len int32 = 1024
  promiscuous bool = false
  err error
  timeout time.Duration = 30 * time.Second
  handle *pcap.Handle
)

func main() {
  handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
  if err != nil { log.Fatal(err) }
  defer handle.Close()

  /*
  var filter string = "tcp and port 80"
  err = handle.SetBPFFilter(filter)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("only capturing TCP port 80 packets.")
  */

  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    fmt.Println(packet)
  }
}