# WeClaw

[中文文档](README_CN.md)

WeChat AI Agent Bridge — connect WeChat to AI coding agents (Claude, Codex, Gemini, Kimi, etc.) via the [iLink](https://www.ilink.wiki) API.

![Preview](preview.png)

## How It Works

```
WeChat User
    │
    ▼
iLink API (long-poll)
    │
    ▼
┌─────────────────────────────────┐
│           WeClaw                │
│                                 │
│  Monitor ──► Handler ──► Agent  │
│                │                │
│                ▼                │
│             Sender              │
└─────────────────────────────────┘
    │                         │
    ▼                         ▼
WeChat Reply            AI Agent Process
                        (ACP / CLI / HTTP)
```

**Agent modes:**

| Mode | How it works | Examples |
|------|-------------|----------|
| ACP  | Long-running subprocess, JSON-RPC over stdio. Fastest — reuses process and sessions. | Claude, Codex, Kimi, Gemini, Cursor, OpenCode, OpenClaw |
| CLI  | Spawns a new process per message. Supports session resume via `--resume`. | Claude (`claude -p`), Codex (`codex exec`) |
| HTTP | OpenAI-compatible chat completions API. | OpenClaw (HTTP fallback) |

Auto-detection picks ACP over CLI when both are available.

## Quick Start

```bash
# Install
go install github.com/fastclaw-ai/weclaw@latest

# Start (first run will prompt QR code login)
weclaw start
```

On first start, WeClaw will prompt a QR code for WeChat login, then auto-detect installed agents and save the config to `~/.weclaw/config.json`. Use `weclaw login` to add additional WeChat accounts.

## Chat Commands

Send these commands as WeChat messages:

### Talk to an agent

```
hello                     # sends to the default agent
/codex write a function   # sends to codex
/cc explain this code     # sends to claude (alias)
```

### Switch default agent

```
/claude       # switch default to claude
/codex        # switch default to codex
/gemini       # switch default to gemini
```

The switch is persisted to config — survives restarts.

### Status

```
/status       # show current agent, type, and model
```

### Aliases

| Alias | Agent |
|-------|-------|
| `/cc` | claude |
| `/cx` | codex |
| `/cs` | cursor |
| `/km` | kimi |
| `/gm` | gemini |
| `/ocd` | opencode |
| `/oc` | openclaw |

## Configuration

Config file: `~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "model": "sonnet"
    },
    "codex": {
      "type": "cli",
      "command": "/usr/local/bin/codex"
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main"
    }
  }
}
```

Environment variables:
- `WECLAW_DEFAULT_AGENT` — override default agent
- `OPENCLAW_GATEWAY_URL` — OpenClaw HTTP fallback endpoint
- `OPENCLAW_GATEWAY_TOKEN` — OpenClaw API token

## Docker

```bash
# Build
docker build -t weclaw .

# Login (interactive — scan QR code)
docker run -it -v ~/.weclaw:/root/.weclaw weclaw login

# Start with HTTP agent
docker run -d --name weclaw \
  -v ~/.weclaw:/root/.weclaw \
  -e OPENCLAW_GATEWAY_URL=https://api.example.com \
  -e OPENCLAW_GATEWAY_TOKEN=sk-xxx \
  weclaw

# View logs
docker logs -f weclaw
```

> Note: ACP and CLI agents require the agent binary inside the container.
> The Docker image ships only WeClaw itself. For ACP/CLI agents, mount
> the binary or build a custom image. HTTP agents work out of the box.

## Development

```bash
# Hot reload
make dev

# Build
go build -o weclaw .

# Run
./weclaw start
```

## License

[MIT](LICENSE)
