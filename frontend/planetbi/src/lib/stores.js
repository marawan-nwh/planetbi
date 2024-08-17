import { writable } from "svelte/store";

export const title = writable("- PlanetBI");

// https://stackoverflow.com/a/71601719/17931895
export const _isDevelopment = import.meta.env.DEV;
