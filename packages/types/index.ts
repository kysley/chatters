export type BTTVUserResponse = {
  id: string;
  bots: unknown[];
  channelEmotes: {
    id: string;
    code: string;
    imageType: string;
    userId: string;
  }[];
  sharedEmotes: {
    id: string;
    code: string;
    imageType: string;
  }[];
};

export type FourPieceState = {
  claps: number;
  user: string;
  emote: string;
};

export enum ChattersEventType {
  "EMOTE" = "EMOTE",
  "FOUR_PIECE" = "FOUR_PIECE",
  "COMBO" = "COMBO",
}

export type ChattersEventPayload = {
  [ChattersEventType.EMOTE]: { name: string; count: number };
  [ChattersEventType.FOUR_PIECE]: FourPieceState | "CLEAR";
  [ChattersEventType.COMBO]: { name: string; count: number } | "CLEAR";
};

export interface ChattersServerEvents {
  [ChattersEventType.EMOTE]: (payload: ChattersEventPayload["EMOTE"]) => void;
  [ChattersEventType.FOUR_PIECE]: (
    payload: ChattersEventPayload["FOUR_PIECE"]
  ) => void;
  [ChattersEventType.COMBO]: (payload: ChattersEventPayload["COMBO"]) => void;
}
