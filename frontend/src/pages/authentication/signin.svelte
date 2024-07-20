<script lang="ts">
  import { link } from "svelte-spa-router";
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

<main class="flex justify-center items-center h-screen bg-[#76abae]">
  <div class="max-w-[500px] w-full p-5 bg-[var(--bg-secondary)] rounded-md shadow-md flex flex-col">
    <h2 class="text-center" id="title">Biscord | Welcome back</h2>
    <p class="text-center text-[#96b1b3]" id="caption">We're excited to see you again!</p>

    <form on:submit={login} class="flex flex-col w-full">
      <!-- email -->
      <label for="email" class="font-bold text-xs mt-5 text-[#76abae]">EMAIL</label>
      <input bind:value={email} id="email" type="email" class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" />

      <!-- password -->
      <label for="password" class="font-bold text-xs mt-5 text-[#76abae]">PASSWORD</label>
      <input bind:value={password} id="password" type="password" class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" />
      <a href="/forgot-password" class="text-left text-xs mt-1 text-[#76abae] no-underline hover:underline">Forgot your password?</a>
      <p class="text-red-500">
        {errorMessage}
      </p>
      <button type="submit" class="bg-[#76abae] text-[var(--bg-secondary)] p-4 font-bold border-none rounded-sm cursor-pointer mt-5">
        {isLoading ? "Loading..." : "Sign in"}
      </button>
    </form>

    <p class="mt-1 text-xs text-left">
      Don't have an account? <a href="/register" use:link class="text-[#76abae] no-underline hover:underline">Register</a>
    </p>
  </div>
</main>