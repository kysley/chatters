import { ChattersEventPayload } from "types";

export function createEmoteEvent(
  name: string,
  count: number
): ChattersEventPayload["EMOTE"] {
  return {
    name,
    count,
  };
}
