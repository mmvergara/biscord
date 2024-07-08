import Signin from "./pages/auth/Signin.svelte";
import Signup from "./pages/auth/Signup.svelte";
const publicRoutes = {
  "/": Signin,
  "/register": Signup,
};

export const routes = {
  ...publicRoutes,
};
