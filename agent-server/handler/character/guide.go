package character

import (
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	"github.com/sirupsen/logrus"
)

type GuideFlag uint64

const (
	None         GuideFlag = 0
	Welcome      GuideFlag = 1
	Minimap      GuideFlag = 2
	WearArmor    GuideFlag = 4
	MonIdentify  GuideFlag = 8
	HowToFight   GuideFlag = 16
	GetItem      GuideFlag = 32
	RecoveryLife GuideFlag = 64
	LevelUp      GuideFlag = 128
	HwanMode     GuideFlag = 256
	HowToParty   GuideFlag = 512
	Condition    GuideFlag = 1024
	HowToPK      GuideFlag = 2048
	JobChoice    GuideFlag = 4096
	Trader       GuideFlag = 8192
	Hunter       GuideFlag = 16384
	Bandit       GuideFlag = 32768
	MovingStep   GuideFlag = 65536
	CarryStep    GuideFlag = 131072
	Criminal     GuideFlag = 262144  //unverified
	Fine         GuideFlag = 524288  //unverified
	Quest        GuideFlag = 1048576 //unverified
	Apprentice   GuideFlag = 2097152
	StreetStall  GuideFlag = 4194304
	GetSkill     GuideFlag = 8388608
	BuyItem      GuideFlag = 16777216
	Action       GuideFlag = 33554432
	Academy      GuideFlag = 67108864
	OpenMarket   GuideFlag = 134217728
)

type GuideHandler struct {
	channel chan server.PacketChannelData
}

func InitGuideHandler() {
	queue := server.PacketManagerInstance.GetQueue(opcode.GuideRequest)
	handler := GuideHandler{channel: queue}
	go handler.Handle()
}

func (h *GuideHandler) Handle() {
	for {
		data := <-h.channel
		// TODO implement real logic
		gflag, err := data.ReadUInt64()
		if err != nil {
			logrus.Panicf("failed to read guide flag")
		}

		p := network.EmptyPacket()
		p.MessageID = opcode.GuideResponse
		p.WriteByte(1)
		p.WriteUInt64(gflag)
		data.Conn.Write(p.ToBytes())
	}
}
