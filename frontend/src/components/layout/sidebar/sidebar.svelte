<script>
  import HomeSidebar from "./home-sidebar/home-sidebar.svelte";
  import { location } from "svelte-spa-router";

  // Subscribe to the location store
  location.subscribe((newLocation) => {
    console.log("Current route:", newLocation);
    // Do something with the new location
  });
</script>

<aside class="bg-black2 w-[240px] fixed left-[72px] h-screen flex flex-col">
  <HomeSidebar />
</aside>
