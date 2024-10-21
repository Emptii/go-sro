package party

import (
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	log "github.com/sirupsen/logrus"
)

type PartyMatchingDeleteHandler struct {
	channel chan server.PacketChannelData
}

func InitPartyMatchingDeleteHandler() {
	handler := PartyMatchingDeleteHandler{channel: server.PacketManagerInstance.GetQueue(opcode.PartyMatchingDeleteRequest)}
	go handler.Handle()
}

func (h *PartyMatchingDeleteHandler) Handle() {
	partyService := service.GetPartyServiceInstance()
	for {
		data := <-h.channel
		partyNumber, err := data.ReadUInt32()
		if err != nil {
			log.Panicln("Failed to read partyNumber")
		}

		partyService.DeletePartyFromMatching(partyNumber, data.UserContext.UniqueID)

	}
}
