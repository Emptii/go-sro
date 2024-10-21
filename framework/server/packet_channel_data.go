package server

import "github.com/Emptii/go-sro/framework/network"

type PacketChannelData struct {
	*Session
	network.Packet
}
