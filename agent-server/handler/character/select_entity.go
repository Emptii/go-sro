package character

import (
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	"github.com/sirupsen/logrus"
)

type SelectEntityHandler struct {
	channel chan server.PacketChannelData
}

func InitSelectEntityHandler() {
	queue := server.PacketManagerInstance.GetQueue(opcode.EntitySelectRequest)
	handler := SelectEntityHandler{channel: queue}
	go handler.Handle()
}

func (h *SelectEntityHandler) Handle() {
	for {
		data := <-h.channel
		// TODO implement real logic
		uniqueId, err := data.ReadUInt32()
		if err != nil {
			logrus.Panicf("failed to read unique id")
		}

		world := service.GetWorldServiceInstance()
		object, err := world.GetObjectByUniqueId(uniqueId)

		p := network.EmptyPacket()
		p.MessageID = opcode.EntitySelectResponse
		if err != nil {
			logrus.Error(err)
			p.WriteByte(0)
			p.WriteByte(0)
			data.Conn.Write(p.ToBytes())
			return
		}
		p.WriteByte(1)
		p.WriteUInt32(uniqueId)
		if object.GetTypeInfo().IsNPCNpc() {
			// TODO: depends to talk options, so ignore for now
			refChar, err := service.GetReferenceDataServiceInstance().GetReferenceCharacter(object.GetRefObjectID())
			if err != nil {
				logrus.Debugf("Selected entity %s is not an interactable npc", object.GetName())
			}

			logrus.Debugf("Selected refChar %s %s", refChar.OrgObjCodeName, refChar.CodeName)

			p.WriteByte(0)

			// the eSRO project uses the following code here:
			// pkt->Write<uint8_t>(tasks.size());
			// for (std::vector<uint8_t>::const_iterator i = tasks.begin(); i != tasks.end(); ++i)
			//	pkt->Write<uint8_t>(*i);
			// see https://github.com/ghostuser846/eSRO/blob/c69c9465c1a61cb880b475315f3cc5b96e78158a/EPL/src/packet_npc.cpp#L161
			// By the way, tasks seem to get loaded on world start and eSRO does it like this:
			// https://github.com/ghostuser846/eSRO/blob/c69c9465c1a61cb880b475315f3cc5b96e78158a/SQL/property_query.cpp#L43
			p.WriteByte(0)
			p.WriteByte(0)
		} else if object.GetTypeInfo().IsNPCMob() {

			refChar, err := service.GetReferenceDataServiceInstance().GetReferenceCharacter(object.GetRefObjectID())
			if err != nil {
				logrus.Debugf("Selected entity %s is not an npc", object.GetName())
			}

			p.WriteByte(1)
			p.WriteUInt32(uint32(refChar.CurrentHealth)) // TODO: Monster HP
			p.WriteByte(1)
			p.WriteByte(5)

		} else if object.GetTypeInfo().IsPlayerCharacter() {
			p.WriteUInt32(0)
			p.WriteByte(0) // TODO: Trader Level
			p.WriteByte(0) // TODO: Hunter Level
			p.WriteByte(0) // TODO: Thief Level
			p.WriteByte(0)
		}
		data.Conn.Write(p.ToBytes())
	}
}
