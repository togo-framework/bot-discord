---
name: bot-discord
description: Run a Discord bot in a togo app — configure BOT_DRIVER=discord + DISCORD_BOT_TOKEN and register handlers with the bot plugin
---

# togo bot-discord

Discord driver for the togo `bot` subsystem.

## Setup

```bash
togo install togo-framework/bot
togo install togo-framework/bot-discord
```

1. Create an app + bot at the [Discord Developer Portal](https://discord.com/developers/applications).
2. **Enable the Message Content Intent** (Bot → Privileged Gateway Intents) — the
   bot can't read command text without it.
3. Invite the bot to your server (`bot` scope).
4. `.env`:
   ```bash
   BOT_DRIVER=discord
   DISCORD_BOT_TOKEN=...
   ```
5. Register handlers with `bot.OnCommand` / `bot.OnMessage` (see the `bot` skill).

## Notes
- `m.Channel` is the Discord channel ID; reply with `Service.Send`.
- The driver ignores its own + other bots' messages (no loops).
- Never commit the token; keep it in `.env`.
