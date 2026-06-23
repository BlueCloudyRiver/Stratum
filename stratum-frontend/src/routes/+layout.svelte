<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/stratum.png';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
{@render children()}
