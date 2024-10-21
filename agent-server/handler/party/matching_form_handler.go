package party

import (
	"github.com/Emptii/go-sro/agent-server/model"
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	log "github.com/sirupsen/logrus"
)

type PartyMatchingFormHandler struct {
	channel chan server.PacketChannelData
}

func InitPartyMatchingFormHandler() {
	handler := PartyMatchingFormHandler{channel: server.PacketManagerInstance.GetQueue(opcode.PartyMatchingFormRequest)}
	go handler.Handle()
}

func (h *PartyMatchingFormHandler) Handle() {
	for {
		data := <-h.channel
		_, err := data.ReadUInt32() // TODO: what is this? Possible Placeholder for partynumber
		if err != nil {
			log.Panicln("Failed to read")
		}

		_, err1 := data.ReadUInt32() // TODO: what is this? Possible Placeholder for partynumber
		if err1 != nil {
			log.Panicln("Failed to read")
		}

		partySetting, err2 := data.ReadByte()
		if err2 != nil {
			log.Panicln("Failed to read partySetting")
		}

		purposeType, err3 := data.ReadByte()
		if err3 != nil {
			log.Panicln("Failed to read purposeType")
		}

		levelMin, err4 := data.ReadByte()
		if err4 != nil {
			log.Panicln("Failed to read levelMin")
		}

		levelMax, err5 := data.ReadByte()
		if err5 != nil {
			log.Panicln("Failed to read levelMax")
		}

		title, err6 := data.ReadString()
		if err6 != nil {
			log.Panicln("Failed to read title")
		}

		party := model.PartyFormRequest{
			MasterJID:         data.UserContext.UserID,
			MasterUniqueID:    data.UserContext.UniqueID,
			MasterName:        data.UserContext.CharName,
			CountryType:       0, // TODO: Figure out
			PartySettingsFlag: model.PartySetting(partySetting),
			PurposeType:       model.PartyPurpose(purposeType),
			LevelMin:          levelMin,
			LevelMax:          levelMax,
			Title:             title,
		}

		partyService := service.GetPartyServiceInstance()
		partyNumber := partyService.FormParty(party)

		// TODO Add check for existing party
		p := network.EmptyPacket()
		p.MessageID = opcode.PartyMatchingFormResponse
		p.WriteByte(1) // TODO possibly result?
		p.WriteUInt32(partyNumber)
		p.WriteUInt32(0)
		p.WriteByte(partySetting)
		p.WriteByte(purposeType)
		p.WriteByte(levelMin)
		p.WriteByte(levelMax)
		p.WriteString(title)
		data.Session.Conn.Write(p.ToBytes())
	}
}
