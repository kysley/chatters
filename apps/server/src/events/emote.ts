import { ChattersEventPayload } from "types";

export function createEmoteEvent(
  payload: ChattersEventPayload["EMOTE"]
): ChattersEventPayload["EMOTE"] {
  return payload;
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
