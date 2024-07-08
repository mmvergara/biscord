// stores/userStore.ts
import { writable } from "svelte/store";
import axios from "axios";
import { API_URL } from "../../config";

export type User = {
  username: string;
  exp: number;
  userid: string;
};

function createUserStore() {
  const { subscribe, set } = writable<User | null>(null);

  // Load user from localStorage on initialization
  if (typeof window !== "undefined") {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      set(JSON.parse(storedUser));
    }
  }

  // Subscribe to changes and update localStorage
  subscribe((value) => {
    if (typeof window !== "undefined") {
      localStorage.setItem("user", JSON.stringify(value));
    }
  });

  return {
    subscribe,
    signIn: async (username: string, password: string) => {
      const res = await axios.post(
        API_URL + "auth/signin",
        { username, password },
        { withCredentials: true },
      );
      set(res.data as User);
    },
    signUp: async (username: string, password: string) => {
      await axios.post(
        API_URL + "auth/signup",
        { username, password },
        { withCredentials: true },
      );
      console.log("Sign up successful");
    },
    signOut: async () => {
      await axios.post(API_URL + "auth/signout", null, {
        withCredentials: true,
      });
      set(null);
    },
  };
}

// Create the store only once
export const userStore = createUserStore();
