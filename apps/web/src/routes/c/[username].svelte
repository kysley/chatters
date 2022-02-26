<script lang="ts">
	import { page } from '$app/stores';
	import Emote from '../../components/Emote.svelte';
	import { Chain } from '../../zeus';

	const chain = Chain('http://localhost:3610/graphql');

	let username = $page.params.username;

	async function getChatter() {
		try {
			const { chatter } = await chain('query')({
				chatter: [
					{ username },
					{
						id: true,
						occurances: {
							uses: true,
							emote: {
								code: true,
								emoteId: true
							}
						}
					}
				]
			});
			return chatter;
		} catch (e) {
			return;
		}
	}
</script>

<h1>{$page.params.username} chatting stats</h1>
{#await getChatter()}
	<p>loading...</p>
{:then data}
	{#each data.occurances as occ}
		<Emote emote={{ code: occ.emote.code, id: occ.emote.emoteId, imageType: 'fu' }} />
		{occ.uses}
	{/each}
{/await}

<svelte:head>
	<title>{$page.params.username}</title>
</svelte:head>
