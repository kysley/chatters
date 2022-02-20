<script lang="ts">
	import { fade } from 'svelte/transition';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventType, ChattersServerEvents, EmoteAndCount, FourPieceState } from 'types';
	import lru from 'tiny-lru';

	import Emote from '../components/Emote.svelte';

	let occuranceLru = lru<EmoteAndCount>(10, 15 * 1000);

	$: keys = [];

	export const prerender = true;

	const socketUrl =
		import.meta.env.MODE === 'production' ? 'https://api.e8y.fun' : 'http://localhost:3610';

	const socket: Socket<ChattersServerEvents> = io(socketUrl, { path: '/chatters/socket.io' });

	$: occurances = [] as EmoteAndCount[];
	socket.on(ChattersEventType.EMOTE, (payload) => {
		// console.log(payload);
		// occurances = [payload, ...occurances];

		const lruItem = occuranceLru.get(payload.emote.id);
		if (lruItem) {
			occuranceLru.set(payload.emote.id, { ...lruItem, count: lruItem.count + payload.count });
		} else {
			occuranceLru.set(payload.emote.id, payload);
		}

		keys = occuranceLru.keys();
	});

	let combo: EmoteAndCount;
	$: combo;
	socket.on(ChattersEventType.COMBO, (payload) => {
		if (payload === 'CLEAR') {
			setTimeout(() => {
				combo = undefined;
			}, 1000 * 3);
			return;
		}
		combo = payload;
	});

	let fourPiece: FourPieceState;
	$: fourPiece;
	socket.on(ChattersEventType.FOUR_PIECE, (payload) => {
		if (payload === 'CLEAR') {
			fourPiece = undefined;
			return;
		}
		fourPiece = payload;
	});
</script>

<h1>Welcome to chatters</h1>
<p>Visit <a href="https://twitch.tv/moonmoon">twitch.tv/moonmoon</a></p>

{#if combo}
	COMBO <Emote emote={combo.emote} /> x {combo.count}
{/if}

{#if fourPiece}
	{fourPiece.emote} combo! nice Clap :) {fourPiece.user}. You have {fourPiece.claps}.
{/if}
<ul>
	{#if keys.length}
		{#each keys as emoteItem}
			{@const emoteLru = occuranceLru.get(emoteItem)}
			{#if emoteLru}
				<div in:fade>
					<Emote emote={emoteLru.emote} /> x {emoteLru.count}
				</div>
			{/if}
		{/each}
	{/if}
</ul>
