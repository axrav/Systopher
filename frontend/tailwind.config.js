/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./app/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      animation: {
        "alert-in": "alert-in 0.6s cubic-bezier(.8,.1,.68,.94)",
        "shake-xy": "shake-xy 1s ease-in-out",
        "error-shake": "error-shake 0.7s ease-in",
        "rotate-shake": "rotate-shake 0.7s ease-in",
      },
      keyframes: {
        "alert-in": {
          "0%": {
            transform: "scale(1, 1)",
            fontWeight: "300",
            color: "rgb(254 226 226)",
          },
          "50%": {
            transform: "scale(1.3, 1.3)",
            fontWeight: "600",
            color: "rgb(153 27 27)",
          },
          "100%": {
            transform: "scale(1, 1)",
            fontWeight: "300",
            color: "rgb(220 38 38)",
          },
        },

        "shake-xy": {
          "0%": {
            transform: "rotate(0)",
            transform: "scale(1)",
            fontWeight: "300",
            color: "rgb(248 113 113)",
          },
          "10%, 30%, 70%, 90%": {
            transform: "rotate(-5deg)",
          },
          "50%": {
            transform: "scale(1.5)",
            fontWeight: "600",
            color: "rgb(220 38 38)",
          },
          "20%, 40%, 60%, 80%": {
            transform: "rotate(5deg)",
          },
          "100%": {
            transform: "scale(1)",
            fontWeight: "300",
            color: "rgb(220 38 38)",
          },
        },
      },
    },
    plugins: [require("tailwind-scrollbar-hide")],
  },
};
