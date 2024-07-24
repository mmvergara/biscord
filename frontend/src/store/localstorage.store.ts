import { writable } from "svelte/store";

export const createLocalStore = <T>(key: string, startValue: T) => {
  const localStorageValue = localStorage.getItem(key);
  const initialValue = localStorageValue
    ? JSON.parse(localStorageValue)
    : startValue;

  const store = writable<T>(initialValue);

  store.subscribe((value) => {
    localStorage.setItem(key, JSON.stringify(value));
  });

  return store;
};
