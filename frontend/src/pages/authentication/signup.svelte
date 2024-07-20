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

<main class="flex justify-center items-center h-screen bg-[#76abae]">
  <form class="max-w-[500px] w-full p-5 bg-[var(--bg-secondary)] rounded-md shadow-md flex flex-col" on:submit={register}>
    <h2 class="text-center" id="title">Biscord | Welcome back</h2>
    <p class="text-center text-[#96b1b3]" id="caption">We're excited to see you again!</p>

    <!-- email -->
    <label for="email" class="font-bold text-xs mt-3 mb-1 text-[#76abae]">EMAIL</label>
    <input bind:value={email} class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" id="email" type="email" />
    
    <!-- displayname -->
    <label for="displayName" class="font-bold text-xs mt-3 mb-1 text-[#76abae]">DISPLAY NAME</label>
    <input bind:value={displayName} class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" id="displayName" type="text" />

    <!-- username -->
    <label for="username" class="font-bold text-xs mt-3 mb-1 text-[#76abae]">USERNAME</label>
    <input bind:value={username} class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" id="username" type="text" />

    <!-- password -->
    <label for="password" class="font-bold text-xs mt-3 mb-1 text-[#76abae]">PASSWORD</label>
    <input bind:value={password} class="bg-[var(--bg-primary)] p-4 font-bold text-[#76abae] outline-none border-none rounded-sm" id="password" type="password" />
    
    <p class="text-red-500">
      {errorMessage}
    </p>
    <button type="submit" class="bg-[#76abae] text-[var(--bg-secondary)] p-4 font-bold border-none rounded-sm cursor-pointer mt-5">
      {isLoading ? "Signing up..." : "Sign Up"}
    </button>

    <p class="mt-1 text-xs text-left">
      Already have an account? <a href="/" use:link class="text-[#76abae] no-underline hover:underline">Login</a>
    </p>
  </form>
</main>