import { Chain } from '../zeus';
import type { Emote } from 'types';

export const socketUrl =
	import.meta.env.MODE === 'production' ? 'https://api.e8y.fun' : 'http://localhost:3610';

export const apiUrl =
	import.meta.env.MODE === 'production'
		? 'https://api.e8y.fun/chatters/graphql'
		: 'http://localhost:3610/chatters/graphql';

export const makeBTTVEmoteUrl = (emote: Emote) => `https://cdn.betterttv.net/emote/${emote.id}/2x`;
export const makeChannelEmoteUrl = (emote: Emote) =>
	`https://static-cdn.jtvnw.net/emoticons/v2/${emote.id}/static/light/3.0`;

export const chain = Chain(apiUrl);

export const makeRandomTransform = (height: number, width: number) => {
	const x = Math.floor(width * 0.75 * Math.random());
	let y = Math.floor(height * 0.75 * Math.random());

	console.log({ y, height });
	if (height - y < 300) {
		console.log('too far down');
		y -= 150;
		console.log({ newY: y });
	}

	return `transform: translate(${x}px, ${y}px);`;
};
