package discord

import (
	"context"
	"os"
	"testing"

	"github.com/togo-framework/bot"
)

func TestDriverRegistered(t *testing.T) {
	found := false
	for _, n := range bot.Drivers() {
		if n == "discord" {
			found = true
		}
	}
	if !found {
		t.Fatal("discord driver not registered with bot")
	}
}

func TestFactoryRequiresToken(t *testing.T) {
	old := os.Getenv("DISCORD_BOT_TOKEN")
	_ = os.Unsetenv("DISCORD_BOT_TOKEN")
	defer os.Setenv("DISCORD_BOT_TOKEN", old)

	b, err := makeDriver(nil, nil)
	if err == nil || b != nil {
		t.Fatalf("expected error without DISCORD_BOT_TOKEN, got (%v,%v)", b, err)
	}
}

// With a (fake) token the factory should build a session without dialing.
func TestFactoryBuildsSession(t *testing.T) {
	old := os.Getenv("DISCORD_BOT_TOKEN")
	os.Setenv("DISCORD_BOT_TOKEN", "fake.token.value")
	defer os.Setenv("DISCORD_BOT_TOKEN", old)

	b, err := makeDriver(nil, func(_ context.Context, _ bot.Message) {})
	if err != nil || b == nil {
		t.Fatalf("expected a driver, got (%v,%v)", b, err)
	}
}
