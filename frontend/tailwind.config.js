// tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        black1: "#212226",
        black2: "#2c2d30",
        black3: "#323338",
        black4: "#35383d",
        blackHover: "#36373c",
        link: "#3498db",
      },
    },
  },
  plugins: [],
};
