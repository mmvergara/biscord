<script>
  import ServerList from "./server-list/server-list.svelte";
  import Sidebar from "./sidebar/sidebar.svelte";
</script>

<main class="bg-black3 h-screen">
  <ServerList />
  <Sidebar />
  <section class="ml-[312px] w-full">
    <slot />
  </section>
</main>

<style>
  main {
    display: flex;
    flex-direction: row;
  }
</style>
