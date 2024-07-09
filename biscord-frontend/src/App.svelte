<script lang="ts">
  import Router from "svelte-spa-router";
  import { userStore } from "./store/auth.store";
  import Home from "./pages/app/home.svelte";
  import NotFound from "./pages/not-found.svelte";
  import Signin from "./pages/auth/Signin.svelte";
  import Signup from "./pages/auth/Signup.svelte";
  import { derived } from "svelte/store";

  const isAuthenticated = derived(userStore, ($userStore) => !!$userStore);
  const publicRoutes = {
    "/": Signin,
    "/register": Signup,
  };

  const privateRoutes = {
    "/": Home,
  };

  $: {
    console.log("isAuthenticated", $isAuthenticated);
  }
</script>

{#if $isAuthenticated}
  <Router routes={privateRoutes} />
{:else}
  <Router routes={publicRoutes} />
{/if}
