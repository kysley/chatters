<script lang="ts">
	import type { Emote } from 'types';
	import { makeBTTVEmoteUrl, makeChannelEmoteUrl } from '../utils';

	export let emote: Emote;

	$: url = emote.code.includes('moon2') ? makeChannelEmoteUrl(emote) : makeBTTVEmoteUrl(emote);
</script>

<img src={url} alt={emote.code} style="height: 100%" />
