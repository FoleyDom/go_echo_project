/** @type {import('tailwindcss').Config} */
module.exports = {
   content: [
      "./views/**/*.{html,js,templ,go}",
      "./views/layout/**/*.{html,js,templ,go}",
      "./views/common/**/*.{html,js,templ,go}",
      "./views/components/**/*.{html,js,templ,go}"
   ],
   theme: {
      extend: {},
   },
   plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};