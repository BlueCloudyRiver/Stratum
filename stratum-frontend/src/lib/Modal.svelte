<script lang="ts">
  import { createEventDispatcher } from "svelte";

  let title = $state("");
  const dispatch = createEventDispatcher();

  const submit = (event: Event) => {
    event.preventDefault();
    if (!title.trim()) return;
    dispatch("create", { title: title.trim() });
    title = "";
  };

  const handleClose = () => dispatch("close");
</script>

<div class="fixed inset-0 bg-black/50 flex items-center justify-center p-4">
  <div class="w-full max-w-md rounded bg-white p-6 shadow-lg">
    <h2 class="text-xl font-semibold mb-4">New Project</h2>
    <form onsubmit={submit}>
      <label class="mb-2 block text-sm font-medium text-slate-700" for="title">
        Project title
      </label>
      <input
        id="title"
        class="w-full rounded border border-slate-300 px-3 py-2 text-sm focus:border-sky-500 focus:outline-none"
        bind:value={title}
        type="text"
        placeholder="Enter project title"
      />
      <div class="mt-4 flex justify-end gap-2">
        <button type="button" class="rounded border px-4 py-2 text-sm" onclick={handleClose}>
          Cancel
        </button>
        <button type="submit" class="rounded bg-sky-600 px-4 py-2 text-sm text-white hover:bg-sky-700">
          Create
        </button>
      </div>
    </form>
  </div>
</div>
