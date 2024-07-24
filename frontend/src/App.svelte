<script lang="ts">
  import Router from "svelte-spa-router";
  import { userStore } from "./store/auth.store";
  import NotFound from "./pages/not-found.svelte";
  import { derived } from "svelte/store";
  import Layout from "./components/layout/layout.svelte";
  import Signup from "./pages/authentication/signup.svelte";
  import Signin from "./pages/authentication/signin.svelte";
  import Home from "./pages/app/channels/me/me.svelte";

  const isAuthenticated = derived(userStore, ($userStore) => !!$userStore);
  const publicRoutes = {
    "/": Signin,
    "/register": Signup,
    "*": NotFound,
  };

  const privateRoutes = {
    "/channels/@me": Home,
    "/channels/:guildId": Home,
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
