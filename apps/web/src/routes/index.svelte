<script lang="ts">
	import { fade, fly } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import io, { Socket } from 'socket.io-client';
	import { ChattersEventPayload, ChattersEventType, ChattersServerEvents } from 'types';

	let arr: ChattersEventPayload['EMOTE'][] = [];

	const name = import.meta.env.MODE;

	const socketUrl =
		import.meta.env.MODE === 'production'
			? 'https://api.e8y.fun/chatters'
			: 'http://localhost:3600';

	const socket: Socket<ChattersServerEvents> = io(socketUrl);

	socket.on(ChattersEventType.EMOTE, (payload) => {
		arr = [payload, ...arr];
	});
</script>

<h1>Welcome to chatters</h1>
<p>Visit <a href="https://twitch.tv/moonmoon">twitch.tv/moonmoon</a></p>

<ul>
	{#each arr as emoteItem (emoteItem)}
		<li in:fade out:fly={{ x: 100 }}>{emoteItem.name} x {emoteItem.count}</li>
	{/each}
</ul>
