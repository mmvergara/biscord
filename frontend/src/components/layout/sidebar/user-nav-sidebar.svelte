<script>
  import Avatar from "../../avatar.svelte";
</script>

<article class="h-[50px] bg-[#21252a] flex p-2 gap-1">
  <div class="w-[50%] flex gap-2 items-center">
    <Avatar
      status="online"
      avatarUrl="https://cdn.discordapp.com/avatars/903528902759182356/75b7fbf95678d5bc740e15a577f671ad.webp?size=80"
    />
    <div class="flex flex-col leading-3">
      <span class="text-sm text-gray-100 text-nowrap">crownie 👑</span>
      <p class="text-xs text-gray-500">invinsible</p>
    </div>
  </div>

  <ul class="flex items-center text-gray-400 gap-2 w-[50%] justify-around px-2">
    <li class="hover:bg-zinc-700 p-1 rounded-sm cursor-pointer">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="lucide lucide-mic"
        ><path d="M12 2a3 3 0 0 0-3 3v7a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3Z" /><path
          d="M19 10v2a7 7 0 0 1-14 0v-2"
        /><line x1="12" x2="12" y1="19" y2="22" /></svg
      >
    </li>
    <li class="hover:bg-zinc-700 p-1 rounded-sm cursor-pointer">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="lucide lucide-headphones"
        ><path
          d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"
        /></svg
      >
    </li>
    <li class="hover:bg-zinc-700 p-1 rounded-sm cursor-pointer">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="lucide lucide-bolt"
        ><path
          d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"
        /><circle cx="12" cy="12" r="4" /></svg
      >
    </li>
  </ul>
</article>
