<script lang="ts">
	import type { Emote } from 'types';

	export let emote: Emote;

	const makeBTTVEmoteUrl = (emote: Emote) => `https://cdn.betterttv.net/emote/${emote.id}/2x`;
</script>

<img src={makeBTTVEmoteUrl(emote)} alt={emote.code} style="height: 45px" />
