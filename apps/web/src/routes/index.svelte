<script lang="ts">
	import { expoIn } from 'svelte/easing';
	import { fade, fly } from 'svelte/transition';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventType, ChattersServerEvents, EmoteAndCount, FourPieceState } from 'types';
	import lru from 'tiny-lru';

	import Emote from '../components/Emote.svelte';
	import { socketUrl } from '../utils';

	let occuranceLru = lru<EmoteAndCount>(10, 15 * 1000);

	$: keys = [];

	export const prerender = true;

	const socket: Socket<ChattersServerEvents> = io(socketUrl, { path: '/chatters/socket.io' });

	socket.on(ChattersEventType.EMOTE, (payload) => {
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

	function pop(node, { duration }) {
		return {
			duration,
			css: (t) => {
				const eased = expoIn(t) + 1;

				return `
					transform: scale(${eased});
			`;
			}
		};
	}
</script>

<h1>Welcome to chatters</h1>
<p>Visit <a href="https://twitch.tv/moonmoon">twitch.tv/moonmoon</a></p>

{#if combo}
	COMBO <Emote emote={combo.emote} /> x {combo.count}
{/if}

{#if fourPiece}
	{fourPiece.emote} combo! nice Clap :) {fourPiece.user}. You have {fourPiece.claps}.
{/if}

{#if keys.length}
	<div style="display: flex; flex-direction: column; gap: 0.5rem;">
		{#each keys as emoteItem}
			{@const emoteLru = occuranceLru.get(emoteItem)}
			{#if emoteLru}
				<div style="display: flex; align-items: flex-end; gap: 0.25rem;">
					{#key emoteLru.count}
						<div in:pop={{ duration: 200 }} style="height: {emoteLru.count * 2 + 50}px;">
							<Emote emote={emoteLru.emote} />
						</div>
						<div in:fly={{ x: 10, duration: 100 }} style="font-size: {0.7 + emoteLru.count / 8}rem">
							{emoteLru.count + 'x'}
						</div>
					{/key}
				</div>
			{/if}
		{/each}
	</div>
{/if}
