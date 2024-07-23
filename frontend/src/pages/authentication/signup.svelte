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

<div class="flex justify-center items-center h-screen bg-black2">
  <div
    class="max-w-[550px] w-full p-8 bg-black3 rounded-md flex flex-col uniform-shadow text-white"
  >
    <h2 class="text-center text-lg font-bold" id="title">
      Biscord | Create an account
    </h2>
    <p class="text-center text-gray-400 font-light" id="caption">
      We're excited to have you join us!
    </p>

    <form on:submit={register} class="flex flex-col w-full">
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

      <!-- displayname -->
      <label for="displayName" class="font-bold text-xs mt-5 mb-1 text-gray-300"
        >DISPLAY NAME</label
      >
      <input
        bind:value={displayName}
        id="displayName"
        type="text"
        class="bg-black1 p-[0.7rem] text-sm outline-none border-none rounded-sm"
      />

      <!-- username -->
      <label for="username" class="font-bold text-xs mt-5 mb-1 text-gray-300"
        >USERNAME</label
      >
      <input
        bind:value={username}
        id="username"
        type="text"
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

      <p class="text-red-500">
        {errorMessage}
      </p>
      <button
        type="submit"
        class="bg-[#76abae] p-[0.7rem] border-none rounded-sm cursor-pointer mt-5 font-bold tracking-wide"
      >
        {isLoading ? "Signing Up" : "Register"}
      </button>
    </form>

    <p class="mt-1 text-xs text-left">
      Already have an account?
      <a href="/" use:link class="text-link no-underline hover:underline"
        >Log In</a
      >
    </p>
  </div>
</div>
