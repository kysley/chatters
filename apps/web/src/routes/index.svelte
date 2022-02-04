<script lang="ts">
	import { fade, fly } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventPayload, ChattersEventType, ChattersServerEvents } from 'types';

	export const prerender = true;

	let arr: ChattersEventPayload['EMOTE'][] = [];

	const socketUrl =
		import.meta.env.MODE === 'production'
			? 'https://api.e8y.fun/chatters/'
			: 'http://localhost:3610';

	const socket: Socket<ChattersServerEvents> = io(socketUrl);

	$: last = null;
	$: combo = 0;
	$: combos = [] as Array<[string, number]>;
	socket.on(ChattersEventType.EMOTE, (payload) => {
		if (last?.name === payload.name) {
			combo += 1;
			// combos = [...combos, [payload.name]]
		} else {
			combo = 1;
		}
		last = payload;
		arr = [payload, ...arr];
	});
</script>

<h1>Welcome to chatters</h1>
<p>Visit <a href="https://twitch.tv/moonmoon">twitch.tv/moonmoon</a></p>

{#if combo !== 1}
	COMBO {combo} x {last?.name}
{/if}
<ul>
	{#each arr as emoteItem (emoteItem)}
		<li in:fade out:fly={{ x: 100 }}>{emoteItem.name} x {emoteItem.count}</li>
	{/each}
</ul>
