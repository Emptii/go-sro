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
		skillId, err := data.ReadUInt32()
		if err != nil {
			logrus.Panicf("failed to read unique id")
		}
		logrus.Info("skillId %s", skillId)

		p := network.EmptyPacket()
		p.MessageID = opcode.EntitySkillIncreaseResponse
		p.WriteByte(0x01)
		p.WriteUInt32(skillId)

		data.Conn.Write(p.ToBytes())
	}
}
