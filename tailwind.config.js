/** @type {import('tailwindcss').Config} */
module.exports = {
   content: [
      "./view/**/*.{html,js,templ,go}",
      "./view/layout/**/*.{html,js,templ,go}",
   ],
   theme: {
      extend: {},
   },
   plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};