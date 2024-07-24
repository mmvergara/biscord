<script>
  import Avatar from "../../../../../components/avatar.svelte";
</script>

<article
  class="h-[60px] cursor-pointer hover:bg-blackStripe rounded-md flex flex-col pb-2 px-2"
>
  <div class="border-t-[1px] w-full mx-auto border-blackStripe"></div>
  <div class="flex justify-between items-center mt-3">
    <div class="flex gap-2 items-center">
      <Avatar
        avatarUrl="https://cdn.discordapp.com/avatars/903528902759182356/75b7fbf95678d5bc740e15a577f671ad.webp?size=80"
        status="online"
      />
      <div>
        <div class="text-white font-bold text-sm">Kenny</div>
        <div class="text-gray-400 text-xs">Playing Valorant</div>
      </div>
    </div>
    <div>
      <button
        class="p-2 bg-black2 rounded-full text-gray-400 hover:text-gray-200"
      >
        <svg
          aria-hidden="true"
          role="img"
          xmlns="http://www.w3.org/2000/svg"
          width="20"
          height="20"
          fill="none"
          viewBox="0 0 24 24"
          ><path
            fill="currentColor"
            d="M12 22a10 10 0 1 0-8.45-4.64c.13.19.11.44-.04.61l-2.06 2.37A1 1 0 0 0 2.2 22H12Z"
            class=""
          ></path></svg
        >
      </button>
      <button
        class="p-2 bg-black2 rounded-full text-gray-400 hover:text-gray-200"
      >
        <svg
          aria-hidden="true"
          role="img"
          xmlns="http://www.w3.org/2000/svg"
          width="20"
          height="20"
          fill="none"
          viewBox="0 0 24 24"
          ><path
            fill="currentColor"
            fill-rule="evenodd"
            d="M10 4a2 2 0 1 0 4 0 2 2 0 0 0-4 0Zm2 10a2 2 0 1 1 0-4 2 2 0 0 1 0 4Zm0 8a2 2 0 1 1 0-4 2 2 0 0 1 0 4Z"
            clip-rule="evenodd"
            class=""
          ></path></svg
        >
      </button>
    </div>
  </div>
</article>
