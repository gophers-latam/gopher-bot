package messages

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/db"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore msg by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	sqlite := db.Respuestas{}
	msg, err := sqlite.GetMsg(m.Content)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, `**No sé, no tengo respuesta para el comando**`)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, msg.Respuesta)
	}
}
