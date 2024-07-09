<script lang="ts">
  import Router from "svelte-spa-router";
  import { userStore } from "./store/auth.store";
  import Home from "./pages/app/home.svelte";
  import NotFound from "./pages/not-found.svelte";
  import Signin from "./pages/auth/Signin.svelte";
  import Signup from "./pages/auth/Signup.svelte";
  import { derived } from "svelte/store";
  import Layout from "./components/layout/layout.svelte";

  const isAuthenticated = derived(userStore, ($userStore) => !!$userStore);
  const publicRoutes = {
    "/": Signin,
    "/register": Signup,
    "*": NotFound,
  };

  const privateRoutes = {
    "/": Home,
    "*": NotFound,
  };

  $: {
    console.log("isAuthenticated", $isAuthenticated);
  }
</script>

{#if $isAuthenticated}
  <Layout>
    <Router routes={privateRoutes} />
  </Layout>
{:else}
  <Router routes={publicRoutes} />
{/if}
