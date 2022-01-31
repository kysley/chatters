import { Server } from "socket.io";
import tmi from "tmi.js";

import { createEmoteEvent } from "./events";
import { emoteSet } from "./utils";

const twitchClient = new tmi.Client({
  options: { debug: true },
  channels: ["moonmoon"],
});

export async function runTmi(io: Server) {
  try {
    await twitchClient.connect();
  } catch (e) {
    throw e;
  }

  twitchClient.on("message", (channel, tags, message, self) => {
    if (self) return;

    const words = message.split(" ");

    const occurances: Record<string, number> = {};

    for (const word of words) {
      if (emoteSet.has(word)) {
        if (occurances[word]) {
          occurances[word]++;
        } else {
          occurances[word] = 1;
        }
      }
    }

    for (const key in occurances) {
      console.log(key, occurances[key]);
      const event = createEmoteEvent(key, occurances[key]);
      io.emit("EMOTE", event);
    }
  });
}
