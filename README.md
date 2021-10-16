# chatters backend

chatters counts the usage of TTV emotes in [moonmoon's](twitch.tv/moonmoon) chat.

chatters is pretty simple. The backend connects to chat via IRC, and broadcasts its findings to any connected clients via websocket. It also records some stats, stored in an SQLite db :)

## Getting Started

To get going on your own,
```bash
> cd chatters
> go run ./ #should open to localhost:81
```

## Deployment
The intended deployment is on a vpc with Docker and nginx already running. I'm sure you can repurpose this for running on something like fly.io

`docker build . -t chatters`

`docker run -dit --name <container_name> -p 8081:81 -v chatters-data:/usr/bin/chatters-data chatters`

This will bind 81 in the docker container to 8081 on the machines localhost. Write a server block on your nginx config like the following

```
server {
  listen 81;
  server_name <your.url.here>;

  location / {
    proxy_pass http://localhost:8081;
  }
}
```
and hopefully it runs :)
