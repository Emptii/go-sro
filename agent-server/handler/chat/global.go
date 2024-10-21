package chat

import (
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
)

func handleGlobalMessage(request MessageRequest, session *server.Session) {
	// TODO: Remove global from players inventory
	p := network.EmptyPacket()
	p.MessageID = opcode.ChatUpdate
	p.WriteByte(request.ChatType)
	p.WriteString(session.UserContext.CharName)
	p.WriteString(request.Message)

	service.GetWorldServiceInstance().BroadcastRaw(p.ToBytes())
}
