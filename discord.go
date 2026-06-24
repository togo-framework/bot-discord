// Package discord is the Discord driver for togo's bot subsystem. Blank-import
// it alongside github.com/togo-framework/bot and set BOT_DRIVER=discord plus
// DISCORD_BOT_TOKEN to run a Discord bot.
//
//	import _ "github.com/togo-framework/bot"
//	import _ "github.com/togo-framework/bot-discord"
package discord

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/togo-framework/bot"
	"github.com/togo-framework/togo"
)

func init() {
	bot.RegisterDriver("discord", makeDriver)
}

func makeDriver(k *togo.Kernel, dispatch func(context.Context, bot.Message)) (bot.Bot, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("bot-discord: DISCORD_BOT_TOKEN is not set")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("bot-discord: %w", err)
	}
	// We need message content to read commands.
	session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsMessageContent
	return &driver{session: session, dispatch: dispatch}, nil
}

type driver struct {
	session  *discordgo.Session
	dispatch func(context.Context, bot.Message)
	botID    string
}

// Start opens the gateway connection and registers the message handler. It blocks
// until ctx is canceled or Stop is called.
func (d *driver) Start(ctx context.Context) error {
	d.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore our own messages to avoid loops.
		if m.Author == nil || m.Author.ID == s.State.User.ID || m.Author.Bot {
			return
		}
		d.dispatch(ctx, bot.Message{
			Channel:  m.ChannelID,
			User:     m.Author.ID,
			Username: m.Author.Username,
			Text:     m.Content,
			Platform: "discord",
			Raw:      map[string]any{"message": m.Message},
		})
	})
	if err := d.session.Open(); err != nil {
		return fmt.Errorf("bot-discord: open gateway: %w", err)
	}
	if d.session.State != nil && d.session.State.User != nil {
		d.botID = d.session.State.User.ID
	}
	<-ctx.Done()
	return d.session.Close()
}

// Stop closes the gateway connection.
func (d *driver) Stop() error {
	if d.session != nil {
		return d.session.Close()
	}
	return nil
}

// Send posts msg to a channel by ID.
func (d *driver) Send(ctx context.Context, channel, msg string) error {
	_, err := d.session.ChannelMessageSend(channel, msg)
	if err != nil {
		return fmt.Errorf("bot-discord: send: %w", err)
	}
	return nil
}
