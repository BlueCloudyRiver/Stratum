<script lang="ts">

    let email="";
    let password="";

    async function register() {
        const response = await fetch("http://localhost:8080/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify ({
                email,
                password,
            })
        });

        const resp = await response.json();

        console.log(resp);
    }
</script>

<h1 class="text-3xl font-bold mb-4">Stratum Register</h1>


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
	onclick={register}
>
	Register
</button>