import Fastify, { FastifyRequest } from "fastify";
import FastifyCors from "fastify-cors";

const cache = new Map();

const fastify = Fastify({ trustProxy: true });

fastify.register(FastifyCors, {
  origin: ["http://localhost:8800", "https://chatters.e8y.fun"],
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

  if (cache.has(uri)) {
    res.send(cache.get(uri));
    return;
  }
});

const start = async () => {
  try {
    await fastify.listen(3600);
  } catch (e) {
    console.error(e);
    process.exit(1);
  }
};

start();
