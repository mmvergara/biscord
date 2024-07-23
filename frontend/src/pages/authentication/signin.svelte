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

<div class="flex justify-center items-center h-screen bg-black2 text-white">
  <div
    class="max-w-[550px] w-full p-8 bg-black3 rounded-md flex flex-col uniform-shadow"
  >
    <h2 class="text-center text-lg font-bold" id="title">
      Biscord | Welcome back
    </h2>
    <p class="text-center text-gray-400 font-light" id="caption">
      We're excited to see you again!
    </p>

    <form on:submit={login} class="flex flex-col w-full">
      <!-- email -->
      <label for="email" class="font-bold text-xs mt-5 mb-1 text-gray-300"
        >EMAIL</label
      >
      <input
        bind:value={email}
        id="email"
        type="email"
        class="bg-black1 p-[0.7rem] text-sm outline-none border-none rounded-sm"
      />

      <!-- password -->
      <label for="password" class="font-bold text-xs mt-5 mb-1 text-gray-300"
        >PASSWORD</label
      >
      <input
        bind:value={password}
        id="password"
        type="password"
        class="bg-black1 p-[0.7rem] text-sm outline-none border-none rounded-sm"
      />
      <a
        href="/forgot-password"
        class="text-left text-xs mt-2 text-link no-underline hover:underline"
        >Forgot your password?</a
      >
      <p class="text-red-500">
        {errorMessage}
      </p>
      <button
        type="submit"
        class="bg-[#76abae] p-[0.7rem] border-none rounded-sm cursor-pointer mt-5 tracking-wide font-bold"
      >
        {isLoading ? "Logging In" : "Log In"}
      </button>
    </form>

    <p class="mt-1 text-xs text-left">
      Don't have an account? <a
        href="/register"
        use:link
        class="text-link no-underline hover:underline">Register</a
      >
    </p>
  </div>
</div>
