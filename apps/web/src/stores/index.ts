import { page } from '$app/stores';
import { chain } from '../utils';

export const createEmoteStore = (code: string) => {
	return {
		subscribe: (cb: Function) => {
			const query = chain('query')({});
		}
	};
};
