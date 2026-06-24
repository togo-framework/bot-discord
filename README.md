<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/bot-discord</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/bot-discord"><img src="https://pkg.go.dev/badge/github.com/togo-framework/bot-discord.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Discord driver for the <a href="https://github.com/togo-framework/bot">togo bot</a> subsystem.</strong></p>
</div>

## Install

```bash
togo install togo-framework/bot
togo install togo-framework/bot-discord
```

<!-- /togo-header -->

The **Discord** driver for togo's [`bot`](https://github.com/togo-framework/bot)
subsystem, built on [discordgo](https://github.com/bwmarrin/discordgo). It opens a
gateway connection and dispatches each message to the bot command/message handlers
you register once with `bot.OnCommand` / `bot.OnMessage`.

## Configure

1. Create an application + bot at the
   [Discord Developer Portal](https://discord.com/developers/applications),
   copy the **Bot Token**.
2. Enable the **Message Content Intent** (Bot → Privileged Gateway Intents).
3. Invite the bot to your server with the `bot` scope.
4. Set env:
   ```bash
   BOT_DRIVER=discord
   DISCORD_BOT_TOKEN=...
   ```

Blank-import the driver next to the base:

```go
import (
	_ "github.com/togo-framework/bot"
	_ "github.com/togo-framework/bot-discord"
)
```

`m.Channel` is the Discord channel ID; `Service.Send(ctx, channel, msg)` posts to
it. The bot ignores its own and other bots' messages to avoid loops.

## License

MIT © togo-framework

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
