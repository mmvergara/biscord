<script>
  import ServerIcon from "./server-icon.svelte";
</script>

<main>
  <nav>
    <ul>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <div class="divider" />
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
      <li>
        <ServerIcon
          imageUrl="https://cdn.discordapp.com/embed/avatars/0.png"
          serverId="123456789"
        />
      </li>
    </ul>
  </nav>
  <slot />
</main>

<style>
  main {
    display: flex;
    flex-direction: row;
  }
  nav {
    display: flex;
    flex-direction: column;
    background-color: var(--bg-primary);
    height: 100vh;
    padding: 0.8em;
    width: fit-content;
  }

  .divider {
    height: 2px;
    background-color: rgb(54, 54, 54);
    padding: 2px 0;
    margin: 10px 0;
    border-radius: 5px;
  }

  ul {
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    overflow-x: hidden;
  }
  ul::-webkit-scrollbar {
    display: none;
  }

  li {
    list-style: none;
  }
</style>
