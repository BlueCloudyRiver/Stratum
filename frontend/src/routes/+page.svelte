<script lang="ts">
	import logo from '$lib/assets/logo.png';
	// Svelte 5 runes ready

	const features = [
		{
			icon: 'link',
			title: 'Linked work',
			body: 'Ideas, tasks and code stay connected — jump from a discussion straight to the commit that shipped it, in one click.',
			large: true
		},
		{
			icon: 'message',
			title: 'Discussions',
			body: 'Conversations live next to the work they\u2019re about.'
		},
		{
			icon: 'layers',
			title: 'Project spaces',
			body: 'One home per project, organized by default.'
		},
		{
			icon: 'zap',
			title: 'Lightweight workflow',
			body: 'No status meetings. No process for its own sake.'
		}
	];

	const steps = [
		{ n: '01', label: 'Create a workspace' },
		{ n: '02', label: 'Start a project' },
		{ n: '03', label: 'Plan the work' },
		{ n: '04', label: 'Collaborate' },
		{ n: '05', label: 'Ship' }
	];
</script>

<svelte:head>
	<title>Stratum</title>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
	<link
		href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&family=IBM+Plex+Mono:wght@400;500&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<style>
	.font-display {
		font-family: 'Inter', system-ui, sans-serif;
	}
	.font-mono {
		font-family: 'IBM Plex Mono', ui-monospace, monospace;
	}
	.noise {
		background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='180' height='180'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='2' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E");
	}
</style>

<div class="font-display relative min-h-screen overflow-hidden bg-black text-white antialiased selection:bg-[#376996] selection:text-white">

	<!-- Grain -->
	<div class="noise pointer-events-none fixed inset-0 z-30 opacity-[0.035] mix-blend-overlay"></div>

	<!-- Ambient gradient field -->
	<div class="pointer-events-none fixed inset-0">
		<div class="absolute left-1/2 top-[-14%] h-[760px] w-[1200px] -translate-x-1/2 rounded-full bg-[#376996] opacity-[0.18] blur-[190px]"></div>
		<div class="absolute bottom-[8%] right-[-8%] h-[420px] w-[420px] rounded-full bg-[#376996] opacity-[0.10] blur-[150px]"></div>
		<div class="absolute left-[-8%] top-[55%] h-[380px] w-[380px] rounded-full bg-[#376996] opacity-[0.08] blur-[140px]"></div>
	</div>

	<!-- Navbar -->
	<nav class="sticky top-0 z-20 border-b border-white/[0.08] bg-black/50 backdrop-blur-xl">
		<div class="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">

			<div class="flex items-center gap-2.5">
				<img src="/lib/logo.png" alt="Stratum" class="h-6 w-6" />
				<span class="text-[15px] font-semibold tracking-tight">Stratum</span>
			</div>

			<div class="flex items-center gap-6 text-sm">
				<a href="/login" class="text-white/50 transition hover:text-white">
					Login
				</a>

				<a
					href="/register"
					class="rounded-full bg-gradient-to-b from-[#4A7DAD] to-[#2C5478] px-4 py-2 text-sm font-medium text-white shadow-[0_0_0_1px_rgba(255,255,255,0.08),0_4px_16px_-4px_rgba(55,105,150,0.6)] transition hover:shadow-[0_0_0_1px_rgba(255,255,255,0.12),0_6px_24px_-4px_rgba(55,105,150,0.9)]">
					Join
				</a>
			</div>

		</div>
	</nav>

	<!-- Hero -->
	<section class="relative mx-auto flex max-w-6xl flex-col items-center px-6 pb-20 pt-24 text-center sm:pt-32">

		<a
			href="#how-it-works"
			class="font-mono group inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/[0.04] px-4 py-1.5 text-[11px] uppercase tracking-[0.12em] text-white/60 transition hover:border-[#376996]/50 hover:text-white">
			<span class="relative flex h-1.5 w-1.5">
				<span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-[#5B93C4] opacity-60"></span>
				<span class="relative inline-flex h-1.5 w-1.5 rounded-full bg-[#5B93C4]"></span>
			</span>
			Now in early access
		</a>

		<h1 class="mt-8 max-w-3xl text-5xl font-semibold leading-[1.05] tracking-tight sm:text-7xl">
			<span class="bg-gradient-to-b from-white to-white/70 bg-clip-text text-transparent">
				Build software without
			</span>
			<br />
			<span class="bg-gradient-to-r from-[#9BC0E4] via-[#5B93C4] to-[#2C5478] bg-clip-text text-transparent">
				losing context.
			</span>
		</h1>

		<p class="mt-7 max-w-lg text-base leading-relaxed text-white/50 sm:text-lg">
			One collaboration platform for modern engineering teams
			planning, building, and shipping together.
		</p>

		<div class="mt-10 flex flex-col items-center gap-4 sm:flex-row">

			<a
				href="/register"
				class="rounded-full bg-gradient-to-b from-[#4A7DAD] to-[#2C5478] px-7 py-3.5 text-sm font-medium text-white shadow-[0_0_0_1px_rgba(255,255,255,0.08),0_8px_30px_-6px_rgba(55,105,150,0.7)] transition hover:shadow-[0_0_0_1px_rgba(255,255,255,0.14),0_10px_40px_-6px_rgba(55,105,150,1)]">
				Get started
			</a>

			<a
				href="#how-it-works"
				class="rounded-full border border-white/10 px-7 py-3.5 text-sm font-medium text-white/70 transition hover:border-white/25 hover:text-white">
				Learn more
			</a>

		</div>

		<!-- Signature visual: a live "linked work" window -->
		<div
			class="relative mx-auto mt-24 w-full max-w-3xl"
			style="perspective: 1400px;">

			<div class="pointer-events-none absolute -inset-x-10 -inset-y-6 rounded-[32px] bg-[#376996] opacity-20 blur-[80px]"></div>

			<div
				class="relative rounded-2xl border border-white/[0.10] bg-gradient-to-b from-white/[0.06] to-white/[0.015] shadow-[0_40px_100px_-20px_rgba(0,0,0,0.8)] backdrop-blur-2xl"
				style="transform: rotateX(7deg); transform-origin: bottom;">

				<!-- window chrome -->
				<div class="flex items-center gap-1.5 border-b border-white/[0.08] px-5 py-3.5">
					<span class="h-2.5 w-2.5 rounded-full bg-white/15"></span>
					<span class="h-2.5 w-2.5 rounded-full bg-white/15"></span>
					<span class="h-2.5 w-2.5 rounded-full bg-white/15"></span>
					<span class="font-mono ml-3 text-[11px] text-white/30">stratum.app/acme/checkout-redesign</span>
				</div>

				<!-- linked panels -->
				<div class="grid grid-cols-1 gap-px bg-white/[0.06] sm:grid-cols-3">

					<div class="bg-[#0A0A0C] px-5 py-5 text-left">
						<p class="font-mono text-[10px] uppercase tracking-[0.14em] text-[#5B93C4]">Idea</p>
						<p class="mt-2 text-sm text-white/80">Simplify checkout to one page</p>
						<div class="mt-4 space-y-2">
							<div class="h-1.5 w-full rounded-full bg-white/10"></div>
							<div class="h-1.5 w-3/4 rounded-full bg-white/10"></div>
						</div>
					</div>

					<div class="bg-[#0A0A0C] px-5 py-5 text-left">
						<p class="font-mono text-[10px] uppercase tracking-[0.14em] text-[#5B93C4]">Task</p>
						<p class="mt-2 text-sm text-white/80">Build address autofill</p>
						<div class="mt-4 flex items-center gap-2">
							<span class="h-4 w-4 rounded-full border-2 border-[#5B93C4]"></span>
							<div class="h-1.5 w-2/3 rounded-full bg-white/10"></div>
						</div>
					</div>

					<div class="bg-[#0A0A0C] px-5 py-5 text-left">
						<p class="font-mono text-[10px] uppercase tracking-[0.14em] text-[#5B93C4]">Code</p>
						<p class="font-mono mt-2 text-sm text-white/80">#4471 fix(checkout)</p>
						<div class="mt-4 space-y-2">
							<div class="h-1.5 w-1/2 rounded-full bg-white/10"></div>
							<div class="h-1.5 w-1/3 rounded-full bg-[#376996]/50"></div>
						</div>
					</div>

				</div>

				<div class="flex items-center justify-center gap-2 border-t border-white/[0.08] px-5 py-3">
					<span class="h-1 w-1 rounded-full bg-[#5B93C4]"></span>
					<span class="font-mono text-[10px] uppercase tracking-[0.14em] text-white/30">Everything linked, automatically</span>
					<span class="h-1 w-1 rounded-full bg-[#5B93C4]"></span>
				</div>

			</div>
		</div>

	</section>

	<!-- The Problem -->
	<section class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-4xl px-6 py-28">

			<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-[#5B93C4]">
				The problem
			</p>

			<h2 class="mt-5 max-w-xl text-3xl font-semibold leading-tight tracking-tight sm:text-4xl">
				Building software is collaborative.
				<span class="text-white/35">The tools aren&rsquo;t.</span>
			</h2>

			<div class="mt-14 grid gap-px overflow-hidden rounded-2xl border border-white/[0.08] bg-white/[0.06] sm:grid-cols-2">
				<div class="bg-black px-6 py-6 transition hover:bg-white/[0.02]">
					<p class="font-mono text-xs text-white/35">01 — Discussions</p>
					<p class="mt-1.5 text-white/85">happen in chat.</p>
				</div>
				<div class="bg-black px-6 py-6 transition hover:bg-white/[0.02]">
					<p class="font-mono text-xs text-white/35">02 — Ideas</p>
					<p class="mt-1.5 text-white/85">stay in documents.</p>
				</div>
				<div class="bg-black px-6 py-6 transition hover:bg-white/[0.02]">
					<p class="font-mono text-xs text-white/35">03 — Tasks</p>
					<p class="mt-1.5 text-white/85">live in project boards.</p>
				</div>
				<div class="bg-black px-6 py-6 transition hover:bg-white/[0.02]">
					<p class="font-mono text-xs text-white/35">04 — Code</p>
					<p class="mt-1.5 text-white/85">lives in GitHub.</p>
				</div>
			</div>

			<p class="mt-10 text-lg text-white/60">
				Every handoff loses context.
			</p>

		</div>
	</section>

	<!-- The Solution -->
	<section class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-4xl px-6 py-28 text-center">

			<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-[#5B93C4]">
				The solution
			</p>

			<h2 class="mt-5 text-3xl font-semibold tracking-tight sm:text-4xl">
				Everything connected.
			</h2>

			<div class="relative mx-auto mt-16 flex max-w-[240px] flex-col items-center">

				<div class="absolute inset-y-0 left-1/2 w-px -translate-x-1/2 bg-gradient-to-b from-[#376996]/0 via-[#376996]/60 to-[#376996]/0"></div>

				{#each ['Idea', 'Discussion', 'Planning', 'Tasks', 'Code', 'Reviews', 'Ship'] as node, i}
					<div class="relative z-10 flex flex-col items-center">
						<div
							class="flex h-10 w-full items-center justify-center rounded-full border text-sm transition
							{i === 6
								? 'border-[#376996]/60 bg-gradient-to-b from-[#4A7DAD] to-[#2C5478] text-white shadow-[0_0_24px_-4px_rgba(55,105,150,0.9)]'
								: 'border-white/10 bg-black text-white/70'}">
							{node}
						</div>
						{#if i < 6}
							<div class="h-8 w-px"></div>
						{/if}
					</div>
				{/each}

			</div>

			<p class="mt-14 text-white/50">
				One continuous execution workflow.
			</p>

		</div>
	</section>

	<!-- Built for small teams -->
	<section class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-4xl px-6 py-28">

			<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-[#5B93C4]">
				Built for small teams
			</p>

			<h2 class="mt-5 max-w-xl text-3xl font-semibold leading-tight tracking-tight sm:text-4xl">
				Not enterprise software.
			</h2>

			<div class="mt-12 flex flex-wrap gap-3">
				{#each ['Product + engineering', 'Startups', 'Open-source teams', 'Remote teams', 'Fast-moving companies'] as tag}
					<span class="flex items-center gap-2 rounded-full border border-white/10 bg-white/[0.03] px-4 py-2 text-sm text-white/70 transition hover:border-[#376996]/40 hover:text-white">
						<span class="text-[#5B93C4]">✓</span>
						{tag}
					</span>
				{/each}
			</div>

			<p class="mt-12 text-lg text-white/60">
				Built for teams that move fast.
			</p>

		</div>
	</section>

	<!-- Features -->
	<section class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-4xl px-6 py-28">

			<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-[#5B93C4]">
				Features
			</p>

			<h2 class="mt-5 text-3xl font-semibold tracking-tight sm:text-4xl">
				Everything in one place.
			</h2>

			<div class="mt-14 grid grid-cols-1 gap-4 sm:grid-cols-3">
				{#each features as f}
					<div class="group relative overflow-hidden rounded-2xl border border-white/[0.08] bg-white/[0.02] p-7 transition hover:border-[#376996]/40 {f.large ? 'sm:col-span-2' : ''}">
						<div class="pointer-events-none absolute -right-10 -top-10 h-32 w-32 rounded-full bg-[#376996] opacity-0 blur-3xl transition group-hover:opacity-20"></div>

						<div class="flex h-9 w-9 items-center justify-center rounded-lg border border-white/10 bg-white/[0.03] text-[#5B93C4]">
							{#if f.icon === 'link'}
								<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></svg>
							{:else if f.icon === 'message'}
								<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>
							{:else if f.icon === 'layers'}
								<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"><polygon points="12 2 2 7 12 12 22 7 12 2"/><polyline points="2 17 12 22 22 17"/><polyline points="2 12 12 17 22 12"/></svg>
							{:else}
								<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>
							{/if}
						</div>

						<h3 class="mt-5 text-base font-medium text-white">{f.title}</h3>
						<p class="mt-2 max-w-sm text-sm leading-relaxed text-white/50">{f.body}</p>
					</div>
				{/each}
			</div>

		</div>
	</section>

	<!-- How it works -->
	<section id="how-it-works" class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-4xl px-6 py-28">

			<div class="text-center">
				<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-[#5B93C4]">
					How it works
				</p>
				<h2 class="mt-5 text-3xl font-semibold tracking-tight sm:text-4xl">
					Five steps. On repeat.
				</h2>
			</div>

			<div class="relative mt-20 hidden sm:block">
				<div class="absolute left-0 right-0 top-5 h-px bg-gradient-to-r from-[#376996]/0 via-[#376996]/50 to-[#376996]/0"></div>
				<div class="grid grid-cols-5 gap-4">
					{#each steps as step}
						<div class="flex flex-col items-center text-center">
							<span class="font-mono relative z-10 flex h-10 w-10 items-center justify-center rounded-full border border-[#376996]/50 bg-black text-xs text-[#8FB6DC]">
								{step.n}
							</span>
							<span class="mt-4 text-sm text-white/70">{step.label}</span>
						</div>
					{/each}
				</div>
			</div>

			<div class="mt-4 flex flex-col items-center gap-4 sm:hidden">
				{#each steps as step}
					<div class="flex items-center gap-4 self-stretch">
						<span class="font-mono flex h-9 w-9 shrink-0 items-center justify-center rounded-full border border-[#376996]/50 bg-black text-xs text-[#8FB6DC]">
							{step.n}
						</span>
						<span class="text-white/80">{step.label}</span>
					</div>
				{/each}
			</div>

			<p class="font-mono mt-14 text-center text-xs uppercase tracking-[0.25em] text-white/30">
				Repeat
			</p>

		</div>
	</section>

	<!-- Final CTA -->
	<section class="relative border-t border-white/[0.08]">
		<div class="mx-auto max-w-3xl px-6 py-32 text-center">

			<h2 class="text-4xl font-semibold leading-tight tracking-tight sm:text-5xl">
				<span class="bg-gradient-to-b from-white to-white/70 bg-clip-text text-transparent">
					Build software without
				</span>
				<br />
				<span class="bg-gradient-to-r from-[#9BC0E4] via-[#5B93C4] to-[#2C5478] bg-clip-text text-transparent">
					losing context.
				</span>
			</h2>

			<p class="mx-auto mt-6 max-w-sm text-white/50">
				Join the first engineering teams using Stratum.
			</p>

			<a
				href="/register"
				class="mt-10 inline-block rounded-full bg-gradient-to-b from-[#4A7DAD] to-[#2C5478] px-8 py-4 text-sm font-medium text-white shadow-[0_0_0_1px_rgba(255,255,255,0.08),0_10px_40px_-6px_rgba(55,105,150,0.8)] transition hover:shadow-[0_0_0_1px_rgba(255,255,255,0.14),0_12px_50px_-6px_rgba(55,105,150,1)]">
				Get early access
			</a>

		</div>
	</section>

	<!-- Footer -->
	<footer class="relative border-t border-white/[0.08]">
		<div class="mx-auto flex max-w-6xl flex-col items-center gap-6 px-6 py-12 sm:flex-row sm:justify-between">

			<div class="flex items-center gap-2.5">
				<img src="{logo}" alt="Stratum" class="h-5 w-5 opacity-70" />
				<span class="text-sm text-white/50">Stratum</span>
			</div>

			<div class="flex items-center gap-6 text-sm text-white/40">
				<a href="/privacy" class="transition hover:text-white">Privacy</a>
				<a href="/terms" class="transition hover:text-white">Terms</a>
				<a href="https://github.com" class="transition hover:text-white">GitHub</a>
				<a href="/contact" class="transition hover:text-white">Contact</a>
			</div>

		</div>
	</footer>

</div>