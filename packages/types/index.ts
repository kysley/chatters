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

export type Emote = Omit<
  BTTVUserResponse["channelEmotes"][number],
  "imageType" | "userId"
>;

export type EmoteAndCount = {
  emote: Emote;
  count: number;
};

export type FourPieceState = {
  claps: number;
  user: string;
  emote: Emote;
};

export enum ChattersEventType {
  "EMOTE" = "EMOTE",
  "FOUR_PIECE" = "FOUR_PIECE",
  "COMBO" = "COMBO",
}

export type ChattersEventPayload = {
  [ChattersEventType.EMOTE]: EmoteAndCount;
  [ChattersEventType.FOUR_PIECE]: FourPieceState | "CLEAR";
  [ChattersEventType.COMBO]: EmoteAndCount | "CLEAR";
};

export interface ChattersServerEvents {
  [ChattersEventType.EMOTE]: (payload: ChattersEventPayload["EMOTE"]) => void;
  [ChattersEventType.FOUR_PIECE]: (
    payload: ChattersEventPayload["FOUR_PIECE"]
  ) => void;
  [ChattersEventType.COMBO]: (payload: ChattersEventPayload["COMBO"]) => void;
}
