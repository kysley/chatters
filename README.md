# borpa.live backend

borpa.live counts the usage of a family of borpa-based emotes in [moonmoon's](twitch.tv/moonmoon) chat.

borpa.live is pretty simple. The backend connects to chat via IRC, and broadcasts its findings to any connected clients via websocket. It also records some stats, which is saved to a local file with bolt!

## Getting Started

To get going on your own,
```bash
> cd borpa-backend
> go build ./ #should open to localhost:8081
```

