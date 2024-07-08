<script lang="ts">
  let email: string = "";
  let password: string = "";

  function login() {
    console.log("Logging in...");
    console.log({
      email,
      password,
    });
    
  }
</script>

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

  <button on:click={login}>Log In</button>
</form>
<p id="auth-option">
  Don't have an account? <a href="/auth/register">Register</a>
</p>

<style>
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
