<script lang="ts">
	import { makeStatsQuery } from '../graphql';

	async function getStats() {
		const { stats } = await makeStatsQuery();
		return stats;
	}
</script>

{#await getStats()}
	<p>loading stats...</p>
{:then data}
	<p><span>{data.chatters}</span> chatters have used {data.occurances} emotes</p>
{/await}
