<script lang="ts">
	import { expoIn } from 'svelte/easing';
	import { fade, fly } from 'svelte/transition';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventType, ChattersServerEvents, EmoteAndCount, FourPieceState } from 'types';
	import lru from 'tiny-lru';

	import Emote from '../components/Emote.svelte';
	import Stats from '../components/Stats.svelte';
	import { socketUrl, makeRandomTransform } from '../utils';

	let height: number, width: number;

	type LruItem = EmoteAndCount & {
		style: string;
	};

	let occuranceLru = lru<LruItem>(10, 15 * 1000);

	$: keys = [];

	export const prerender = true;

	const socket: Socket<ChattersServerEvents> = io(socketUrl, { path: '/chatters/socket.io' });

	socket.on(ChattersEventType.EMOTE, (payload) => {
		const lruItem = occuranceLru.get(payload.emote.id);
		if (lruItem) {
			occuranceLru.set(payload.emote.id, { ...lruItem, count: lruItem.count + payload.count });
		} else {
			occuranceLru.set(payload.emote.id, { ...payload, style: makeRandomTransform(height, width) });
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

<svelte:window bind:innerHeight={height} bind:innerWidth={width} />
<main class="wrapper">
	<h1>Welcome to chatters</h1>
	<p>Visit <a href="https://twitch.tv/moonmoon">twitch.tv/moonmoon</a></p>
	<Stats />

	<div style="height: 50px;">
		{#if combo}
			COMBO <Emote emote={combo.emote} /> x {combo.count}
		{/if}
	</div>

	{#if fourPiece}
		{fourPiece.emote} combo! nice Clap :) {fourPiece.user}. You have {fourPiece.claps}.
	{/if}

	{#if keys.length}
		<!-- <div class="wrapper"> -->
		{#each keys as emoteItem (emoteItem)}
			{@const emoteLru = occuranceLru.get(emoteItem)}
			{#if emoteLru}
				<div style="position: absolute; {emoteLru.style}">
					{#key emoteLru.count}
						<a href={`/e/${emoteLru.emote.code}`}>
							<div in:pop={{ duration: 200 }} style="height: {emoteLru.count * 2 + 90}px;">
								<Emote emote={emoteLru.emote} />
							</div>
						</a>
						<!-- <div in:fly={{ x: 10, duration: 100 }} style="font-size: {0.7 + emoteLru.count / 8}rem"> -->
						<div
							in:fly={{ x: 10, duration: 100 }}
							style="font-size: 1rem; display: flex; justify-content: space-between;"
						>
							<span>
								x{emoteLru.count}
							</span>
							{#if combo?.emote.code === emoteLru.emote.code}
								<span>
									COMBO x{combo.count}
								</span>
							{/if}
						</div>
					{/key}
				</div>
			{/if}
		{/each}
		<!-- </div> -->
	{/if}
</main>

<!-- <button on:click={addCombo}>add combo</button> -->
<style>
	:global(body) {
		margin: 0;
	}
	:global(html) {
		height: 100vh;
		width: 100vw;
		overflow: hidden;
	}

	h1 {
		margin: 0;
	}
	.wrapper {
		height: 100vh;
		width: 100vw;
	}
</style>
