<script>
    // Importing assets
    import hej from '$lib/assets/hej.jpg';
    import stratum from '$lib/assets/stratum.png';
</script>

<div 
    class="relative min-h-screen bg-cover bg-center bg-no-repeat font-mono antialiased"
    style="background-image: url('{hej}');"
>
    <div class="absolute inset-0 bg-black/75 backdrop-blur-[2px] z-0"></div>

    <nav class="relative z-10 border-b border-white/10 bg-black/40 backdrop-blur-md">
        <div class="max-w-7xl mx-auto h-16 px-6 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div 
                    class="w-6 h-6 bg-contain bg-center bg-no-repeat"
                    style="background-image: url('{stratum}');"
                ></div>
                <span class="text-white font-bold text-xs tracking-widest uppercase font-sans">
                    Stratum
                </span>
            </div>
            
            <div class="flex items-center gap-5 font-sans">
                <a 
                    href="/login" 
                    class="text-xs font-bold uppercase tracking-wider text-gray-400 hover:text-white transition cursor-pointer"
                >
                    Sign In
                </a>

                <a 
                    href="/register" 
                    class="bg-[#158539] text-black text-xs font-bold uppercase tracking-wider px-5 py-2.5 rounded hover:bg-[#00E05B] transition shadow-lg shadow-[#00FF66]/5"
                >
                    Sign Up
                </a>
            </div>
        </div>
    </nav>

    <main class="relative z-10 max-w-7xl mx-auto px-6 pt-16 pb-24 grid lg:grid-cols-12 gap-8 items-center min-h-[calc(100vh-64px)]">
        <div class="lg:col-span-7 space-y-6">
            <div class="inline-block bg-white/5 border border-white/10 text-gray-300 text-[10px] uppercase tracking-widest px-3 py-1 rounded">
                Business Collaboration Platform
            </div>
            
            <h1 class="text-4xl sm:text-6xl font-black text-white font-sans tracking-tight leading-[1.05]">
                Business collaboration, simplified.<br />
                <span class="text-[#158539]">Zero friction.</span>
            </h1>
            
            <p class="text-sm text-gray-300 max-w-xl font-sans leading-relaxed">
                Manage projects, collaborate with your team, organize clients, and streamline workflows from a single intuitive workspace.
            </p>
            
            <div class="pt-2 flex items-center gap-4 text-xs font-bold uppercase">
                <a 
                    href="/register" 
                    class="bg-[#158539] text-black px-6 py-3 rounded hover:bg-[#00E05B] transition shadow-lg shadow-[#00FF66]/10 text-center"
                >
                    Get Started Free
                </a>
                <a href="/contact" class="bg-black/40 text-white border border-white/10 px-6 py-3 rounded hover:bg-black/60 transition backdrop-blur-sm cursor-pointer">
                    Learn more
                </a>
            </div>
        </div>
        
        <div class="lg:col-span-5"></div>
    </main>
</div>