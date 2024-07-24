<script>
  import HomeIcon from "./home-icon.svelte";
  import ServerIcon from "./server-icon.svelte";
</script>

<nav class="flex flex-col fixed bg-black1 h-screen w-[72px] pt-2">
  <ul
    class="flex flex-col overflow-y-auto overflow-x-hidden [&::-webkit-scrollbar]:hidden pt-1 transition-all"
  >
    <HomeIcon imageUrl="https://cdn.discordapp.com/embed/avatars/0.png" />
    <div class="h-0.5 bg-[#363636] mb-2 rounded-md w-[32px] mx-auto py-0.3" />
    {#each Array(13) as _}
      <ServerIcon
        imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
        guildId="123456789"
      />
    {/each}
  </ul>
</nav>
