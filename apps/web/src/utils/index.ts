import { Chain } from '../zeus';

export const socketUrl =
	import.meta.env.MODE === 'production' ? 'https://api.e8y.fun' : 'http://localhost:3610';

export const apiUrl =
	import.meta.env.MODE === 'production'
		? 'https://api.e8y.fun/chatters/graphql'
		: 'http://localhost:3610/graphql';

export const chain = Chain(apiUrl);
