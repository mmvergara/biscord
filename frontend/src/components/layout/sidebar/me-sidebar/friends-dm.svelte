<script>
  import { link } from "svelte-spa-router";
  import Avatar from "../../../avatar.svelte";
</script>

  <li class="hover:bg-blackHover rounded-md text-gray-400">
    <a
      href="/app/home/friends"
      class="flex group items-center gap-4 p-2 pl-3 py-[0.4rem]"
      use:link
    >
      <Avatar
        status="online"
        avatarUrl="https://cdn.discordapp.com/avatars/903528902759182356/75b7fbf95678d5bc740e15a577f671ad.webp?size=80"
      />
      <span class="text-sm mr-auto"> xCrownie </span>
      <span class="group-hover:opacity-100 opacity-0">
        <svg
          x="0"
          y="0"
          role="img"
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          fill="none"
          viewBox="0 0 24 24"
          ><path
            fill="currentColor"
            d="M13 6a1 1 0 1 0-2 0v5H6a1 1 0 1 0 0 2h5v5a1 1 0 1 0 2 0v-5h5a1 1 0 1 0 0-2h-5V6Z"
            class=""
          ></path></svg
        >
      </span>
    </a>
  </li>