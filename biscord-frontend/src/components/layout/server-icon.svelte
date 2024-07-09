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
    border-radius: 50%;
    margin-bottom: 4px;
  }

  .server-icon {
    width: 48px;
    height: 48px;
    object-fit: cover;
    border-radius: 50%;
    transition: all 0.3s ease-in-out;
  }

  .server-icon:hover {
    cursor: pointer;
    border-radius: 30%;
  }

  .server-icon:active {
    transform: scale(0.95);
  }
</style>
