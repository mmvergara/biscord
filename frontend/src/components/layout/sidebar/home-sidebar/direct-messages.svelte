<script>
  import FriendsDm from "./friends-dm.svelte";
</script>

<section class="text-xs text-gray-400">
  <article class="flex justify-between px-3 pt-3 hover:text-gray-200">
    <span> DIRECT MESSAGES </span>
    <svg
      x="0"
      y="0"
      class="privateChannelRecipientsInviteButtonIcon_c47fa9 icon_fc4f04"
      aria-hidden="true"
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
  </article>
  <ul class="p-2 flex gap-1 flex-col">
    <FriendsDm />
    <FriendsDm />
    <FriendsDm />
    <FriendsDm />
    <FriendsDm />
    <FriendsDm />
  </ul>
</section>
