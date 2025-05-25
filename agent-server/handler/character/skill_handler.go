package character

import (
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	"github.com/sirupsen/logrus"
)

type SkillHandler struct {
	channel chan server.PacketChannelData
}

func InitSkillHandler() {
	queue := server.PacketManagerInstance.GetQueue(opcode.EntitySkillIncreaseRequest)
	handler := SkillHandler{channel: queue}
	go handler.Handle()
}

func (h *SkillHandler) Handle() {
	for {
		data := <-h.channel
		// TODO implement real logic
		masteryId, err := data.ReadUInt8()
		logrus.Info("masteryId %s", masteryId)

		if err != nil {
			logrus.Panicf("failed to read unique id")
		}
		p := network.EmptyPacket()
		p.MessageID = opcode.EntitySkillIncreaseResponse
	}
}
