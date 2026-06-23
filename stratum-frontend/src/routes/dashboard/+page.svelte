<script lang="ts">
  import { onMount } from "svelte";
  import Modal from "$lib/Modal.svelte";

  type Project = {
    id: number;
    title: string;
  };

  let projects = $state<Project[]>([]);
  let show = $state(false);

  async function getProjects() {
    const token = localStorage.getItem("token");
    if (!token) return;

    const response = await fetch("http://localhost:8080/projects", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });

    if (!response.ok) {
      console.error("Failed to fetch projects");
      return;
    }

    projects = await response.json();
  }

  onMount(() => {
    getProjects();
  });

  const open = () => (show = true);
  const close = () => (show = false);

  async function create(event: CustomEvent<{ title: string }>) {
    const p = await (await fetch("localhost:8080/project", {
      method: "POST",
      body: JSON.stringify(event.detail),
      headers: { "Content-Type": "application/json" }
    })).json();
    projects = [p, ...projects];
    show = false;
  }
</script>

<div class="flex items-center flex-row p-4 mb-2">
  <h1 class="text-3xl font-bold">Stratum</h1>

  <button class="ml-auto px-4 py-2 border border-gray-400 rounded hover:bg-gray-100" onclick={open}>
    New Project
  </button>
</div>

<hr class="border-gray-400">

<div class="border-1 border-gray-400 rounded m-5 h-120 p-4 mt-7">
  <h1 class="text-xl font-medium pb-3">Projects</h1>
  <hr class="border-gray-400">
  <div class="grid grid-cols-3 gap-4">
  {#each projects as p}
    <div class="border-gray-400 p-4 m-1">
      {p.title}
    </div>
  {/each}
  </div>
  {#if show}
    <Modal on:create={create} on:close={close}/>
  {/if}
</div>


