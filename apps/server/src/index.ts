import Fastify, { FastifyRequest } from "fastify";
import FastifyCors from "fastify-cors";
import got from "got";
import mercurius from "mercurius";
import { Server } from "socket.io";

import { BTTVUserResponse } from "types";
import { schema } from "./schema";
import { runTmi } from "./tmi";
import { prisma, emoteMap } from "./utils";

export const fastify = Fastify({
  trustProxy: true,
});

fastify.register(FastifyCors, {
  origin: ["http://localhost:3000", "https://chatters.e8y.fun"],
});

fastify.register(mercurius, {
  schema,
  // context: (req) => req.ctx,
  graphiql: true,
});

let io: Server;

fastify.get("/health", (_req, res) => {
  res.code(200).send({ statusCode: 200, status: "ok" });
});

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
      path: "/chatters/socket.io",
    });

    io.on("connection", () => console.log("socket connection"));

    const { channelEmotes, sharedEmotes } = await got
      .get("https://api.betterttv.net/3/cached/users/twitch/121059319")
      .json<BTTVUserResponse>();

    const BTTVEmotes = [...channelEmotes, ...sharedEmotes];

    for (const emote of BTTVEmotes) {
      emoteMap.set(emote.code, emote);
    }

    await prisma.emote.createMany({
      skipDuplicates: true,
      data: BTTVEmotes.map(({ code, id }) => ({
        code,
        emoteId: id,
      })),
    });

    await runTmi(io);
  } catch (e) {
    console.error(e);
    process.exit(1);
  }
};

start();
