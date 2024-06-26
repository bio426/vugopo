/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./src/**/*.{vue,ts}",
        "./index.html",
    ],
    theme: { extend: {} },
    plugins: [require("daisyui")],
    daisyui: { themes: ["forest"] }
}

