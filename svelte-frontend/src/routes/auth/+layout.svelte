<main>
  <div class="auth-container">
    <slot />
  </div>
</main>

<style>
  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #76abae;
  }

  .auth-container {
    max-width: 500px;
    width: 100%;
    padding: 20px;
    background-color: var(--bg-secondary);
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
  }
</style>
