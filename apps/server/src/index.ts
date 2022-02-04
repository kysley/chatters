import Fastify, { FastifyRequest } from "fastify";
import FastifyCors from "fastify-cors";
import got from "got";
import { Server } from "socket.io";

import { BTTVUserResponse } from "types";
import { runTmi } from "./tmi";
import { emoteSet } from "./utils";

export const fastify = Fastify({
  trustProxy: true,
});

fastify.register(FastifyCors, {
  origin: ["http://localhost:3000", "https://chatters.e8y.fun"],
});

let io: Server;

type PeerManagerGetRequest = FastifyRequest<{
  Querystring: {
    search: string;
  };
}>;
fastify.get("/", async (req: PeerManagerGetRequest, res) => {
  const { search } = req.query;
  if (!search) return res.status(404).send();

  const uri = encodeURI(search.toLowerCase());

  // if (cache.has(uri)) {
  //   res.send(cache.get(uri));
  //   return;
  // }
});

const start = async () => {
  try {
    await fastify.listen(3610);
    io = new Server(fastify.server, {
      cors: {
        origin: ["http://localhost:3000", "https://chatters.e8y.fun"],
      },
    });

    io.on("connection", () => console.log("socket connection"));

    const { channelEmotes, sharedEmotes } = await got
      .get("https://api.betterttv.net/3/cached/users/twitch/121059319")
      .json<BTTVUserResponse>();

    for (const emote of channelEmotes) {
      emoteSet.add(emote.code);
    }

    for (const emote of sharedEmotes) {
      emoteSet.add(emote.code);
    }

    await runTmi(io);
  } catch (e) {
    console.error(e);
    process.exit(1);
  }
};

start();
