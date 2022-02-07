<script lang="ts">
	import { fade } from 'svelte/transition';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventType, ChattersServerEvents, EmoteAndCount, FourPieceState } from 'types';
	import Emote from '../components/Emote.svelte';

	export const prerender = true;

	const socketUrl =
		import.meta.env.MODE === 'production' ? 'https://api.e8y.fun' : 'http://localhost:3610';

	const socket: Socket<ChattersServerEvents> = io(socketUrl, { path: '/chatters/socket.io' });

	$: occurances = [] as EmoteAndCount[];
	socket.on(ChattersEventType.EMOTE, (payload) => {
		console.log(payload);
		occurances = [payload, ...occurances];
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
	COMBO {combo.emote.code} x {combo.count}
{/if}

{#if fourPiece}
	{fourPiece.emote} combo! nice Clap :) {fourPiece.user}. You have {fourPiece.claps}.
{/if}
<ul>
	{#each occurances as emoteItem (emoteItem)}
		<div in:fade>
			<Emote emote={emoteItem.emote} /> x {emoteItem.count}
		</div>
	{/each}
</ul>
