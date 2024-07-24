<script lang="ts">
  import GuildSidebar from "./guild-sidebar/guild-sidebar.svelte";
  import HomeSidebar from "./me-sidebar/me-sidebar.svelte";
  import { location, params } from "svelte-spa-router";

  // Subscribe to the location store
  location.subscribe((newLocation) => {
    console.log("Current route:", newLocation);
    // Do something with the new location
  });
  let guildId: null | string;
  params.subscribe((newParams) => {
    if (newParams?.guildId) {
      guildId = newParams.guildId;
    } else {
      guildId = null;
    }
  });

  $: console.log(guildId);
</script>

<aside class="bg-black2 w-[240px] fixed left-[72px] h-screen flex flex-col">
  {#if guildId}
    <GuildSidebar {guildId} />
  {:else}
    <HomeSidebar />
  {/if}
</aside>
