# chatters backend

chatters counts the usage of TTV emotes in [moonmoon's](twitch.tv/moonmoon) chat.

chatters is pretty simple. The backend connects to chat via IRC, and broadcasts its findings to any connected clients via websocket. It also records some stats, stored in an SQLite db :)

## Getting Started

To get going on your own,
```bash
> cd chatters
> go run ./    # localhost:8082
```

## Deployment
The intended deployment is on a vpc with Docker and nginx already running. I'm sure you can repurpose this for running on something like fly.io

`docker build . -t chatters`

`docker run -dit --name chattrs -p 8082:8082 -v chatters-data:/usr/bin/chatters-data chatters`

This will bind 8082 in the docker container to 8082 on your machines localhost (-p <nginx proxy port>:<port used in app>). Write a server block on your nginx config like the following

```
server {
        listen 80;
        server_name chatters.e8y.fun;

        location / {
                proxy_pass http://localhost:8082;
        }
}
```
and hopefully it runs :)
