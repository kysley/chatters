<script lang="ts">
	import { page } from '$app/stores';
	import { chain } from '../../utils';

	let code = $page.params.code;

	async function getEmoteUses() {
		try {
			const { uses } = await chain('query')({
				uses: [{ code }, true]
			});
			return uses;
		} catch (e) {
			return;
		}
	}
</script>

<h1>{$page.params.code} emote stats</h1>
{#await getEmoteUses()}
	<p>loading {code}...</p>
{:then data}
	{code} has been used {data} time(s). wow!
{/await}

<svelte:head>
	<title>{$page.params.code}</title>
</svelte:head>
