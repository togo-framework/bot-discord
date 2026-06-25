# bot-discord — docs

**Discord bot.** Discord driver via discordgo — slash/message commands → the registry.

## Install

```bash
togo install togo-framework/bot-discord
```

Registers on the [`bot`](https://github.com/togo-framework/bot) base; select it with **BOT_DRIVER (or bot.provider)**, then use **`togo bot`**.

## Interface

`Bot` — `Start`/`Stop`/`Send`, plus a command/handler registry (`OnCommand`/`OnMessage`) so any plugin can add bot commands.

## Configuration

| Env var | Description |
|---|---|
| `DISCORD_BOT_TOKEN` | Discord bot token (required). |

## Usage & notes

Connects with the bot token, routes message/slash commands to the registry, and sends messages/embeds.

## Example

```bash
togo bot:send '#general' 'Deployed!'
togo bot:ask 'summarize the latest release'
```

## Links

- [discordgo](https://github.com/bwmarrin/discordgo)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/bot-discord)
