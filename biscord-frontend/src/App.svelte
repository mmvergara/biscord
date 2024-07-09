<script lang="ts">
  import Router, { push } from "svelte-spa-router";
  import { userStore } from "./store/auth.store";
  import Signin from "./pages/auth/signin.svelte";
  import Signup from "./pages/auth/signup.svelte";
  import Home from "./pages/app/home.svelte";
  import NotFound from "./pages/not-found.svelte";

  let isAuthenticated = false;
  userStore.subscribe((value) => {
    isAuthenticated = !!value;
  });

  const publicRoutes = {
    "/": Signin,
    "/register": Signup,
  };

  const privateRoutes = {
    "/": Home,
  };

  $: {
    console.log("isAuthenticated", isAuthenticated);
  }
</script>

{#if isAuthenticated}
  <Router routes={privateRoutes} />
{:else}
  <Router routes={publicRoutes} />
{/if}
