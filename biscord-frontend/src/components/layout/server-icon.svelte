<!-- discord server icon sample -->
<script lang="ts">
  import { link } from "svelte-spa-router";
  export let imageUrl = "https://cdn.discordapp.com/embed/avatars/0.png";
  export let serverId = "123456789";
</script>

<a href={`/server/${serverId}`} class="server-icon-link" use:link>
  <img src={imageUrl} alt="discord server icon" class="server-icon" />
</a>

<style>
  .server-icon-link {
    display: block;
    width: 100%;
    height: 100%;
  }

  .server-icon {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 50%;
  }

  .server-icon:hover {
    cursor: pointer;
  }

  .server-icon:active {
    transform: scale(0.95);
  }
</style>
