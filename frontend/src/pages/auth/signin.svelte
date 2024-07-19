<script lang="ts">
  import { link, replace } from "svelte-spa-router";
  import { userStore } from "../../store/auth.store";

  let email: string = "";
  let password: string = "";

  let isLoading: boolean = false;
  let errorMessage: string = "";

  const login = async (e: SubmitEvent) => {
    e.preventDefault();
    isLoading = true;
    const { error } = await userStore.signIn({ email, password });
    isLoading = false;
    if (error) {
      errorMessage = error;
      return;
    }
  };
</script>

<main>
  <form class="auth-container" on:submit={login}>
    <h2 id="title">Biscord | Welcome back</h2>
    <p id="caption">We're excited to see you again!</p>

    <form on:submit={login}>
      <!-- email -->
      <label for="email">EMAIL</label>
      <input bind:value={email} id="email" type="email" />

      <!-- password -->
      <label for="password">PASSWORD</label>
      <input bind:value={password} id="password" type="password" />
      <a href="/forgot-password" id="forgot-password">Forgot your password?</a>
      <p class="error-text">
        {errorMessage}
      </p>
      <button type="submit">
        {isLoading ? "Loading..." : "Sign in"}
      </button>
    </form>

    <p id="auth-option">
      Don't have an account? <a href="/register" use:link>Register</a>
    </p>
  </form>
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

  #title {
    text-align: center;
  }
  #caption {
    text-align: center;
    color: #96b1b3;
  }

  form {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  label {
    font-weight: bold;
    font-size: 0.775rem;
    margin: 0.125rem 0px;
    margin-top: 20px;
    color: #76abae;
  }

  input {
    background-color: var(--bg-primary);
    padding: 1em;
    font-weight: bold;
    color: #76abae;
    outline: none;
    border: none;
    border-radius: 0.125rem;
  }

  #forgot-password {
    text-align: left;
    font-size: 0.75rem;
    margin-top: 5px;
    color: #76abae;
    text-decoration: none;
  }
  #forgot-password:hover {
    text-decoration: underline;
  }

  button {
    background-color: #76abae;
    color: var(--bg-secondary);
    padding: 1em;
    font-weight: bold;
    border: none;
    border-radius: 0.125rem;
    cursor: pointer;
    margin-top: 20px;
  }

  #auth-option {
    margin-top: 5px;
    font-size: 0.75rem;
    text-align: left;
  }
  #auth-option a {
    color: #76abae;
    text-decoration: none;
  }
  #auth-option a:hover {
    text-decoration: underline;
  }
</style>
