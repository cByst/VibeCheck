package vibecheckhandlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// AttachHandlers will attach all vibe check bot handlers to the discord session
func AttachHandlers(discordSession *discordgo.Session) {
	discordSession.AddHandler(commandHandler)
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages written by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check message for command prefix to determine if the message is relevant to the bot
	if strings.HasPrefix(strings.ToLower(m.Content), "!vibecheck ") {
		// check to make sure command is in correct format or dm user that issues the command and tell him the format was wrong and print ussage also delete original mmessage

		//lookup users vibe and return
		log.Infof("")

	} else {
		// get sentimate and store sentimate
	}
}
