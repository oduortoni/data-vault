/*
 * file: go-dv/www/static/js/pages/main/home.js
 * description: This file is used to render the home view.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */
const Home = () => {
    window.app.innerHTML = `
        <section class="hero">
            Home Page
        </section>
    `;

    // Attach event listener to button after rendering
    // document.getElementById("about").onclick = () => window.router.navigate("/about");
};

export default Home;