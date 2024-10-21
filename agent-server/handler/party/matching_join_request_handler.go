package party

import (
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	log "github.com/sirupsen/logrus"
)

type PartyMatchingJoinRequestHandler struct {
	channel chan server.PacketChannelData
}

func InitPartyMatchingJoinRequestHandler() {
	handler := PartyMatchingJoinRequestHandler{channel: server.PacketManagerInstance.GetQueue(opcode.PartyMatchingJoinRequest)}
	go handler.Handle()
}

func (h *PartyMatchingJoinRequestHandler) Handle() {
	partyService := service.GetPartyServiceInstance()
	for {
		data := <-h.channel
		partyNumber, err := data.ReadUInt32()
		if err != nil {
			log.Panicln("Failed to read party number")
		}

		partyService.JoinFormedParty(partyNumber, data.UserContext.UniqueID)
	}
}
