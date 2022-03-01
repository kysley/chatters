import { chain } from '../utils';

export function makeStatsQuery() {
	return chain('query')({
		stats: {
			chatters: true,
			occurances: true
		}
	});
}
