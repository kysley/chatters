<script lang="ts">
	import { fade, fly } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import io, { Socket } from 'socket.io-client';
	import {
		ChattersEventPayload,
		ChattersEventType,
		ChattersServerEvents,
		FourPieceState
	} from 'types';

	export const prerender = true;

	const socketUrl =
		import.meta.env.MODE === 'production' ? 'https://api.e8y.fun' : 'http://localhost:3610';

	const socket: Socket<ChattersServerEvents> = io(socketUrl, { path: '/chatters/socket.io' });

	$: occurances = [] as ChattersEventPayload['EMOTE'][];
	socket.on(ChattersEventType.EMOTE, (payload) => {
		occurances = [payload, ...occurances];
	});

	let combo: { name: string; count: number };
	$: combo;
	socket.on(ChattersEventType.COMBO, (payload) => {
		if (payload === 'CLEAR') {
			combo = undefined;
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
	COMBO {combo.name} x {combo.count}
{/if}

{#if fourPiece}
	{fourPiece.emote} combo! nice Clap :) {fourPiece.user}. You have {fourPiece.claps}.
{/if}
<ul>
	{#each occurances as emoteItem (emoteItem)}
		<li in:fade out:fly={{ x: 100 }}>{emoteItem.name} x {emoteItem.count}</li>
	{/each}
</ul>
