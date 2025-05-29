package character

import (
	"github.com/Emptii/go-sro/agent-server/service"
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

		world := service.GetWorldServiceInstance()
		player, err := world.GetPlayerByUniqueId(data.UserContext.UniqueID)
		if err != nil {
			logrus.Error("failed to get player while handeling MasteryIncreaseRequest")
			return
		}

		skillId, err := data.ReadUInt32()
		if err != nil {
			logrus.Panicf("failed to read unique skillId")
		}
		logrus.Info("skillId %s", skillId)

		player.Skills = append(player.Skills, skillId)

		// Update UpdateCharacterSP
		sendCharacterSPUpdate(data, 100)

		// Update skills
		p := network.EmptyPacket()
		p.MessageID = opcode.EntitySkillIncreaseResponse
		p.WriteByte(0x01)
		p.WriteUInt32(skillId)

		data.Conn.Write(p.ToBytes())
	}
}

func sendCharacterSPUpdate(data server.PacketChannelData, spAmount uint32) {
	p := network.EmptyPacket()
	p.MessageID = opcode.EntityUpdatePoints
	p.WriteByte(uint8(2)) // Assuming this is a fixed point type identifier
	p.WriteUInt32(spAmount)
	p.WriteByte(uint8(0)) // Assuming this is a fixed unknown byte
	data.Conn.Write(p.ToBytes())
}
