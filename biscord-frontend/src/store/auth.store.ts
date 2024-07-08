// stores/userStore.ts
import { writable } from "svelte/store";
import axios from "axios";
import { API_URL } from "../config";
import { api } from "../lib/axios";

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

    if (!storedUser) return;
    const parsedUser = JSON.parse(storedUser);

    // Check if token is expired
    if (parsedUser.exp * 1000 < Date.now()) {
      localStorage.removeItem("user");
      return;
    }
    set(JSON.parse(storedUser));
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
      const res = await api.post(API_URL + "auth/signin", {
        username,
        password,
      });
      console.log(res.data);
      console.log(res.data);
      console.log(res.data);
      console.log(res.data);
      set(res.data as User);
    },
    signUp: async (username: string, password: string) => {
      await api.post(API_URL + "auth/signup", { username, password });
      console.log("Sign up successful");
    },
    signOut: async () => {
      await api.post(API_URL + "auth/signout", null);
      set(null);
    },
  };
}

// Create the store only once
export const userStore = createUserStore();
