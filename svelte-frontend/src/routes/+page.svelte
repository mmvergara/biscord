<script lang="ts">
  let email: string = "";
  let password: string = "";
</script>

<main>
  <div class="auth-container">
    <h2 id="title">Biscord | Welcome back</h2>
    <p id="caption">We're excited to see you again!</p>
    <label for="email">EMAIL</label>
    <input
      bind:value={email}
      id="email"
      type="email"
      placeholder="JohnDoe@gmail.com"
    />
    <label for="password">PASSWORD</label>
    <input
      bind:value={password}
      id="password"
      type="password"
      placeholder="********"
    />
    <a href="/forgot-password" id="forgot-password">Forgot your password?</a>
    <button>Log In</button>
    <p id="auth-option">
      Don't have an account? <a href="/signup">Sign up</a>
    </p>
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

  #title {
    text-align: center;
  }
  #caption {
    text-align: center;
    color: #96b1b3;
  }

  label {
    font-weight: bold;
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
