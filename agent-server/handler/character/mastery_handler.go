package character

import (
	"github.com/Emptii/go-sro/agent-server/service"
	"github.com/Emptii/go-sro/framework/network"
	"github.com/Emptii/go-sro/framework/network/opcode"
	"github.com/Emptii/go-sro/framework/server"
	"github.com/sirupsen/logrus"
)

type MasteryHandler struct {
	channel chan server.PacketChannelData
}

func InitMasteryHandler() {
	logrus.SetLevel(logrus.DebugLevel)
	queue := server.PacketManagerInstance.GetQueue(opcode.EntityMasteryIncreaseRequest)
	handler := MasteryHandler{channel: queue}
	go handler.Handle()
}

func WriteMasteryOrSkill(p *network.Packet, masteryId uint32, masteryLvl byte) {
	p.WriteByte(1)
	p.WriteUInt32(masteryId)
	p.WriteByte(masteryLvl)
}

func (h *MasteryHandler) Handle() {
	for {
		logrus.Info("Handeling EntityMasteryIncreaseRequest")
		data := <-h.channel

		world := service.GetWorldServiceInstance()
		player, err := world.GetPlayerByUniqueId(data.UserContext.UniqueID)
		if err != nil {
			logrus.Error("failed to get player while handeling MasteryIncreaseRequest")
			return
		}
		masteryId, err := data.ReadUInt32()
		masteryLevelStepSize, err := data.ReadByte()
		logrus.Infof("masteryId %s masteryLevel %s , %s", masteryId, masteryLevelStepSize)

		if err != nil {
			logrus.Panicf("failed to read unique id")
		}
		p := network.EmptyPacket()
		p.MessageID = opcode.EntityMasteryIncreaseResponse

		player.Masteries[masteryId]++

		WriteMasteryOrSkill(&p, masteryId, player.Masteries[masteryId])
		//     srv_pkt::MasteryInc(pkt,i->second);
		//     m_connection->Send(pkt);
		//     pkt.reset(new OPacket);
		//     srv_pkt::PlayerStats(pkt,player);
		//     m_connection->Send(pkt);
		//     boost::shared_ptr<Party> party = player->get_party();
		//     return MSG_SUCCESS;

		data.Conn.Write(p.ToBytes())

		player.SendStatsUpdate()
		// int StateGame::OnMasteryInc (IPacket &packet)
		// {
		//     uint32_t mastID = packet.Read<uint32_t>();
		//
		//     uint8_t mastlv = packet.Read<uint8_t>();
		//
		//     if (!packet.EndOfStream())
		//         return MSG_ERROR_SIZE;
		//
		//     if (mastlv != 1)
		//         return MSG_ERROR_ARG;
		//
		//     boost::shared_ptr<Player> player = static_cast<srv::Connection*>(m_connection)->GetPlayer();
		//
		//     if (!player)
		//         return MSG_ERROR;
		//
		//     MasteryTree &tree = player->get_mastery_tree();
		//
		//     MasteryTree::iterator i = tree.find(mastID);
		//
		//     if (i == tree.end())
		//         return MSG_ERROR_ORDER;
		//
		//     boost::shared_ptr<OPacket> pkt(new OPacket);
		//
		//     if (tree.GetMasteryTotal() + 1 > tree.GetMasteryLimit())
		//     {
		//         srv_pkt::MasteryInc(pkt,srv_pkt::MASTERY_ERROR_LIMIT);
		//         m_connection->Send(pkt);
		//         return MSG_SUCCESS;
		//     }
		//
		//     if (i->second.Level() == m_server->GetWorld()->GetMaxCap())
		//     {
		// #if defined DEBUG
		//         syslog(LOG_INFO,"Mastery capped.");
		// #endif
		//         return MSG_SUCCESS;
		//     }
		//
		//     int32_t reqsp = 0;
		//
		//     if (i->second.Level())
		//         reqsp = m_server->GetWorld()->GetRequiredSP(i->second.Level());
		//
		//     if (reqsp == -1)
		//     {
		// #if defined DEBUG
		//         syslog(LOG_INFO,"Required SP for Level %i not found.",i->second.Level()+1);
		// #endif
		//         return MSG_ERROR;
		//     }
		//
		//     if (!player->spend_sp(reqsp))
		//     {
		// #if defined DEBUG
		//         syslog(LOG_INFO,"OnMasteryInc() - Insuficient SP");
		// #endif
		//         return MSG_ERROR_ORDER;
		//     }
		//
		//     ++i->second;
		//
		//     srv_pkt::MasteryInc(pkt,i->second);
		//     m_connection->Send(pkt);
		//
		//     pkt.reset(new OPacket);
		//
		//     srv_pkt::PlayerStats(pkt,player);
		//     m_connection->Send(pkt);
		//
		//     boost::shared_ptr<Party> party = player->get_party();
		//
		//     if (party)
		//     {
		//         std::pair<Mastery,Mastery> masts = tree.GetHighestMastery();
		//         party->SendMasteryNotification(player->GetUniqueID(),masts.first.ID(),masts.second.ID());
		//     }
		//
		//     DB::MASTERY::Update query;
		//     query(m_server->DBConnection(),player->ID(),i->second);
		//
		//     return MSG_SUCCESS;
		// }

	}
}
