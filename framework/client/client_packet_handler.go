package client

import "github.com/Emptii/go-sro/framework/network"

type PacketHandler interface {
	Handle(packet network.Packet)
}
