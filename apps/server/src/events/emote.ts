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

export function createFPEvent(
  state: ChattersEventPayload["FOUR_PIECE"]
): ChattersEventPayload["FOUR_PIECE"] {
  return state;
}

export function createComboEvent(
  payload: ChattersEventPayload["COMBO"]
): ChattersEventPayload["COMBO"] {
  return payload;
}
