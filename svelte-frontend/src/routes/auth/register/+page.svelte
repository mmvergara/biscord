<script lang="ts">
  let email: string = "";
  let password: string = "";
  let displayName: string = "";
  let username: string = "";
</script>

<h2 id="title">Biscord | Welcome back</h2>
<p id="caption">We're excited to see you again!</p>

<!-- email -->
<label for="email">EMAIL</label>
<input bind:value={email} id="email" type="email" />
<!-- displayname -->
<label for="displayName">DISPLAY NAME</label>
<input bind:value={displayName} id="displayName" type="text" />

<!-- username -->
<label for="username">USERNAME</label>
<input bind:value={username} id="username" type="text" />

<!-- password -->
<label for="password">PASSWORD</label>
<input bind:value={password} id="password" type="password" />
<button>Log In</button>
<p id="auth-option">
  Don't have an account? <a href="/signup">Sign up</a>
</p>

<style>
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

  input {
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
