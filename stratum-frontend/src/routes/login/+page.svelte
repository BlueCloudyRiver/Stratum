<script lang="ts">
    let email = "";
    let password = "";

    async function login() {
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                email,
                password,
            })

        });
        
        const resp = await response.json();
        if (!response.ok) {
            console.log("Login failed");
            return;
        }
        localStorage.setItem("token", resp.token);
        const token = localStorage.getItem("token");

        if (!token) {
            window.location.href = "/login";
        }

        window.location.href = "/dashboard";
    }
</script>

<h1 class="text-3xl font-bold mb-4">Stratum Login</h1>

<input
	class="border rounded p-2 w-full mb-3"
	type="email"
	placeholder="Email"
	bind:value={email}
/>

<input
	class="border rounded p-2 w-full mb-3"
	type="password"
	placeholder="Password"
	bind:value={password}
/>

<button
	class="border rounded px-4 py-2"
	on:click={login}
>
	Login
</button>