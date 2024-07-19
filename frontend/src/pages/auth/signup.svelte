<script lang="ts">
  import { link } from "svelte-spa-router";
  import { userStore } from "../../store/auth.store";
  import { push } from "svelte-spa-router";

  let username: string = "";
  let email: string = "";
  let displayName: string = "";
  let password: string = "";
  let errorMessage: string = "";
  let isLoading: boolean = false;

  const register = async (e: SubmitEvent) => {
    e.preventDefault();
    errorMessage = "";
    isLoading = true;
    const { error: err } = await userStore.signUp({
      username,
      email,
      displayName,
      password,
    });
    isLoading = false;
    if (err) {
      console.log(err);
      errorMessage = err;
      return;
    }
    push("/");
  };
</script>

<main>
  <form class="auth-container" on:submit={register}>
    <h2 id="title">Biscord | Welcome back</h2>
    <p id="caption">We're excited to see you again!</p>

    <!-- email -->
    <label for="email">EMAIL</label>
    <input bind:value={email} class="text-input" id="email" type="email" />
    <!-- displayname -->
    <label for="displayName">DISPLAY NAME</label>
    <input
      bind:value={displayName}
      class="text-input"
      id="displayName"
      type="text"
    />

    <!-- username -->
    <label for="username">USERNAME</label>
    <input bind:value={username} class="text-input" id="username" type="text" />

    <!-- password -->
    <label for="password">PASSWORD</label>
    <input
      bind:value={password}
      class="text-input"
      id="password"
      type="password"
    />
    <p class="error-text">
      {errorMessage}
    </p>
    <button type="submit">
      {isLoading ? "Signing up..." : "Sign Up"}
    </button>

    <p id="auth-option">
      Already have an account? <a href="/" use:link>Login</a>
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

  label {
    font-weight: bold;
    font-size: 0.775rem;
    margin: 0.125rem 0px;
    margin-top: 20px;
    color: #76abae;
  }

  .text-input {
    background-color: var(--bg-primary);
    padding: 1em;
    font-weight: bold;
    color: #76abae;
    outline: none;
    border: none;
    border-radius: 0.125rem;
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
