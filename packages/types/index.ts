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

export enum ChattersEventType {
  "EMOTE" = "EMOTE",
}

export type ChattersEventPayload = {
  [ChattersEventType.EMOTE]: {
    name: string;
    count: number;
  };
};

export interface ChattersServerEvents {
  [ChattersEventType.EMOTE]: (payload: ChattersEventPayload["EMOTE"]) => void;
}
