<script lang="ts">
  import UserNavSidebar from "../user-nav-sidebar.svelte";
  import Channel from "./channel.svelte";

  export let guildId: string;
</script>

<div>
  <button
    class=" h-[48px] w-full hover:bg-blackHover flex px-4 items-center justify-between text-sm border-b-black1 text-white border-b-[1.9px]"
  >
    <div class="flex gap-2 items-center justify-center">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="lucide lucide-university"
        ><circle cx="12" cy="10" r="1" /><path
          d="M22 20V8h-4l-6-4-6 4H2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2"
        /><path d="M6 17v.01" /><path d="M6 13v.01" /><path
          d="M18 17v.01"
        /><path d="M18 13v.01" /><path
          d="M14 22v-5a2 2 0 0 0-2-2a2 2 0 0 0-2 2v5"
        /></svg
      >
      <p class="mt-[2px]">Server Name</p>
    </div>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="16"
      height="16"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      class="lucide lucide-chevron-down"><path d="m6 9 6 6 6-6" /></svg
    >
  </button>
</div>
<section class="text-xs text-gray-400 grow-1 flex-1 gap-2 overflow-y-scroll">
  <Channel />
  <Channel />
  <Channel />
  <Channel />
  <Channel />
  <Channel />
</section>
<UserNavSidebar />
