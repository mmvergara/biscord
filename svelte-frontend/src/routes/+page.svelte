
<script lang="ts">
  let email:string = '';
</script>


<main>
<div class="auth-container">
<h1>Welcome back to Biscord</h1>
<p>
  We're excited to see you again!
</p>


<input bind:value={email} type="text">
</div>
  
</main>
<style>


main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #76ABAE;
}

.auth-container {
  width: 400px;
  padding: 20px;
  background-color: var(--bg-secondary);
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}


</style>