<script lang="ts">
  import { link } from "svelte-spa-router";
  export let imageUrl = "https://cdn.discordapp.com/embed/avatars/0.png";
</script>

<li class="server-item">
  <a href={`/home`} class="server-link" use:link>
    <div class="server-indicator"></div>
    <img src={imageUrl} alt="discord server icon" class="server-icon" />
  </a>
</li>

<style>
  .server-item {
    margin-bottom: 8px;
  }

  .server-link {
    display: flex;
    align-items: center;
    width: 100%;
    height: 100%;
  }

  .server-indicator {
    height: 10px;
    width: 5px;
    border-top-right-radius: 4px;
    border-bottom-right-radius: 4px;
    background-color: white;
    margin-right: 7px;
    transition: height 0.3s ease-in-out;
  }

  .server-link:hover .server-indicator {
    height: 35px;
  }

  .server-icon {
    width: 48px;
    height: 48px;
    object-fit: cover;
    border-radius: 50%;
    transition: border-radius 0.3s ease-in-out;
  }

  .server-icon:hover {
    border-radius: 30%;
  }
</style>
