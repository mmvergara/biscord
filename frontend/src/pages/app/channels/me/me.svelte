<script lang="ts">
  import ActiveNow from "./components/active-now.svelte";
  import All from "./components/all.svelte";
  import Blocked from "./components/blocked.svelte";
  import Online from "./components/online.svelte";
  import Pending from "./components/pending.svelte";

  let page: "ONLINE" | "ALL" | "PENDING" | "BLOCKED" = "ONLINE";

  $: console.log(page);
</script>

<div class="h-full">
  <section
    class="px-1 h-[48px] flex items-center border-b-black1 border-b-[1.9px] text-gray-400"
  >
    <div class="flex items-center gap-2 p-2 pl-3 py-[0.6rem]">
      <svg
        role="img"
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        fill="none"
        viewBox="0 0 24 24"
        ><path
          fill="currentColor"
          d="M13 10a4 4 0 1 0 0-8 4 4 0 0 0 0 8Z"
          class=""
        ></path><path
          fill="currentColor"
          d="M3 5v-.75C3 3.56 3.56 3 4.25 3s1.24.56 1.33 1.25C6.12 8.65 9.46 12 13 12h1a8 8 0 0 1 8 8 2 2 0 0 1-2 2 .21.21 0 0 1-.2-.15 7.65 7.65 0 0 0-1.32-2.3c-.15-.2-.42-.06-.39.17l.25 2c.02.15-.1.28-.25.28H9a2 2 0 0 1-2-2v-2.22c0-1.57-.67-3.05-1.53-4.37A15.85 15.85 0 0 1 3 5Z"
          class=""
        ></path></svg
      >
      <span class="text-gray-200 text-sm font-semibold">Friends</span>
    </div>
    <div class="w-[1px] h-[25px] mx-3 bg-blackDivider"></div>
    <ul class="flex list-none gap-4 text-sm tracking-wide text-gray-100">
      <li>
        <button
          on:click={() => (page = "ONLINE")}
          class={`
        ${page === "ONLINE" ? "bg-zinc-700" : "text-gray-300"} p-1 px-3 rounded-[0.2rem]`}
          >Online</button
        >
      </li>
      <li>
        <button
          on:click={() => (page = "ALL")}
          class={`${
            page === "ALL" ? "bg-zinc-700" : "text-gray-300"
          }  p-1 px-3 rounded-[0.2rem]`}>All</button
        >
      </li>
      <li>
        <button
          on:click={() => (page = "PENDING")}
          class={`${page === "PENDING" ? "bg-zinc-700" : "text-gray-300"} p-1 px-3 rounded-[0.2rem]`}
        >
          Pending</button
        >
      </li>
      <li>
        <button
          on:click={() => (page = "BLOCKED")}
          class={`${page === "BLOCKED" ? "bg-zinc-700" : "text-gray-300"} p-1 px-3 rounded-[0.2rem]`}
        >
          Blocked</button
        >
      </li>
      <li>
        <button class="bg-green-700 text-white p-1 px-2 rounded-[0.2rem]">
          Add Friend</button
        >
      </li>
    </ul>
  </section>
  <!-- if -->
  <div class="flex h-[calc(100%-48px)]">
    <section class="grow p-4 px-8">
      {#if page === "ONLINE"}
        <Online />
      {:else if page === "ALL"}
        <All />
      {:else if page === "PENDING"}
        <Pending />
      {:else if page === "BLOCKED"}
        <Blocked />
      {/if}
    </section>

    <section
      class="active-now-tab w-[405px] border-l-[1px] border-blackStripe p-4 h-full"
    >
      <ActiveNow />
    </section>
  </div>
</div>

<style>
  @media (max-width: 1260px) {
    .active-now-tab {
      display: none;
    }
  }
</style>
