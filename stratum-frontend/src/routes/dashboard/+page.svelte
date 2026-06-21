<script lang="ts">
  import { onMount } from "svelte";
  type Project = {
    id: number,
    title: string,
  }
  let projects: Project[] = [];

  async function getProjects() {
    const token = localStorage.getItem("token");

    const response = await fetch("http://localhost:8080/projects", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });

    const data = await response.json();
    projects = data;
  }

  onMount(() => {
    getProjects();
  });
</script>

<h1 class="text-3xl font-bold mb-3 p-4">Stratum</h1>


<div class="border-1 border-black-200 m-5 h-120 p-4">
  <h1 class="text-xl font-medium pb-3">Projects</h1>
  <hr>
  <div class="grid grid-cols-3 gap-4">
  {#each projects as project}
    <div>
      {project.title}
    </div>
  {/each}
  </div>
</div>

