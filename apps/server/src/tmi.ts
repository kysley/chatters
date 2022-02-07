import { Server } from "socket.io";
import tmi from "tmi.js";
import { FourPieceState } from "types";

import { createComboEvent, createEmoteEvent, createFPEvent } from "./events";
import { emoteMap, isFourPiece } from "./utils";

const twitchClient = new tmi.Client({
  options: { debug: true },
  channels: ["moonmoon"],
});

export async function runTmi(io: Server) {
  await twitchClient.connect();

  const controller = new MessageController({ server: io });

  twitchClient.on("message", controller.process);
}

class MessageController {
  countFourPieceApplause = false;
  fpCountDuration = 1000 * 5;
  fpState?: FourPieceState;

  combo = 0;
  comboEmote?: string;

  prevMsg = "";

  io: Server;

  constructor({ server }: { server: Server }) {
    this.io = server;
  }

  ioEmit = (type: string, event: unknown) => {
    this.io.emit(type, event);
  };

  process = (
    channel: string,
    tags: tmi.ChatUserstate,
    message: string,
    self: boolean
  ) => {
    if (self) return;

    const words = message.split(" ");

    const fourPieceEmote = isFourPiece(this.prevMsg, message);
    if (fourPieceEmote) {
      this.countFourPieceApplause = true;
      this.fpState = {
        claps: 0,
        emote: emoteMap.get(fourPieceEmote)!,
        user: tags.username || "???",
      };

      this.ioEmit("FOUR_PIECE", createFPEvent(this.fpState));

      setTimeout(() => {
        this.countFourPieceApplause = false;
        this.fpState = undefined;
        this.ioEmit("FOUR_PIECE", createFPEvent("CLEAR"));
      }, this.fpCountDuration);
    }
    if (this.countFourPieceApplause) {
      if (message.toLowerCase().includes("clap") && this.fpState) {
        this.fpState.claps += 1;
      }
    }

    const occurances: Record<string, number> = {};
    for (const word of words) {
      if (emoteMap.has(word)) {
        if (occurances[word]) {
          occurances[word]++;
        } else {
          occurances[word] = 1;
        }
      }
    }

    const occuranceKeys = Object.keys(occurances);

    for (const key of occuranceKeys) {
      const event = createEmoteEvent({
        emote: emoteMap.get(key)!,
        count: occurances[key],
      });
      this.ioEmit("EMOTE", event);
    }

    // Exact same emote, or message containing only the exact same emote
    const isCombo =
      this.prevMsg === message ||
      (occuranceKeys.length === 1 && occuranceKeys[0] === this?.comboEmote);
    if (isCombo) {
      const repeatedEmote =
        this.prevMsg === message ? message : occuranceKeys[0];
      if (emoteMap.has(repeatedEmote)) {
        this.combo++;
        this.comboEmote = repeatedEmote;
        if (this.combo > 1) {
          this.ioEmit(
            "COMBO",
            createComboEvent({
              count: this.combo,
              emote: emoteMap.get(repeatedEmote)!,
            })
          );
        }
      }
      // Emote/message does not count towards
    } else if (!isCombo && this.combo > 1) {
      this.ioEmit("COMBO", createComboEvent("CLEAR"));
      this.combo = 1;
      this.comboEmote = undefined;
    } else {
      this.combo = 1;
      this.comboEmote = undefined;
    }

    this.prevMsg = message;
  };
}
