// stores/userStore.ts
import { writable } from "svelte/store";
import { API_URL } from "../config";
import { post } from "../lib/axios";
import type { SignInForm, SignUpForm } from "../types/formTypes";

export type UserData = {
  sub: string;
  username: string;
  displayName: string;
  exp: number;
};

function createUserStore() {
  const { subscribe, set } = writable<UserData | null>(null);

  // Load user from localStorage on initialization
  if (typeof window !== "undefined") {
    const storedUser = localStorage.getItem("user");

    if (storedUser) {
      const parsedUser = JSON.parse(storedUser);
      if (parsedUser) {
        if (parsedUser.exp * 1000 < Date.now()) {
          localStorage.removeItem("user");
        }
        set(JSON.parse(storedUser));
      }
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
    signIn: async (signInData: SignInForm) => {
      const res = await post<SignInForm, UserData>(
        API_URL + "auth/sign-in",
        signInData,
      );
      if (res.data) set(res.data);
      return res;
    },
    signUp: async (signUpData: SignUpForm) => {
      return await post(API_URL + "auth/sign-up", signUpData);
    },
    signOut: async () => {
      await post(API_URL + "auth/sign-out", null);
      set(null);
    },
  };
}

// Create the store only once
export const userStore = createUserStore();
