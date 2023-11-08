/** @type {import('tailwindcss').Config} */
export default {
    content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        extend: {
            fontSize: {
                xxs: "0.6rem",
            },
            fontFamily: {
                'custom' : ['Nunito', 'sans-serif']
            }
        },
    },
    plugins: [],
};
