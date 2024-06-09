<aside>
  <nav>
  
  </nav>
</aside>

<main>
  <slot/>
</main>


<style>
  aside {
    width: 72px;
    height: 100vh;
    background-color: var(--bg-primary);
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
  }

  nav {
    padding: 20px;
  }

  main {
    margin-left: 72px;
  }
</style>