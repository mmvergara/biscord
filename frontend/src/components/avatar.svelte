<script lang="ts">
  import type { Status } from "../types/status";

  export let status: Status;
  export let avatarUrl: string;

  $: statusColor =
    status == "dnd"
      ? "bg-red-500"
      : status == "online"
        ? "bg-green-500"
        : status == "idle"
          ? "bg-yellow-500"
          : status == "invisible"
            ? "bg-gray-500"
            : "bg-gray-500";
</script>

<div
  class="relative flex h-8 w-8 items-end justify-end rounded-full text-white"
>
  <img
    src={avatarUrl
      ? avatarUrl
      : "https://cdn.discordapp.com/avatars/903528902759182356/75b7fbf95678d5bc740e15a577f671ad.webp?size=80"}
    alt=" "
    class="absolute rounded-full"
    aria-hidden="true"
    width="40px"
    height="40px"
  />
  <div
    class="h-[12px] w-[12px] rounded-full border-[2px] border-black2 z-[200] absolute"
  >
    <div class={statusColor + " w-full h-full rounded-full"}></div>
  </div>
</div>
