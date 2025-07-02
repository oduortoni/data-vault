/*
 * file: go-dv/www/static/js/pages/main/home.js
 * description: Renders the main landing page with a hero section.
 * author: toni
 * date: 2025-06-28
 * version: 1.1.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */
const Home = () => {
    window.app.innerHTML = `
        <section class="hero-section">
            <div class="hero-content">
                <h1 class="hero-title">Your Data, Your Rules.</h1>
                <p class="hero-subtitle">
                    Data Vault is a secure and flexible platform for creating custom forms, collecting information, and managing it all in one place. You design the structure, you control the access.
                </p>
                <div class="hero-features">
                    <div class="feature-item">
                        <h3>Build Custom Forms</h3>
                        <p>Design forms and tables to capture the exact data you need, without any technical skills.</p>
                    </div>
                    <div class="feature-item">
                        <h3>Secure & Private</h3>
                        <p>Your data is protected with robust security. You decide who can see and manage your information.</p>
                    </div>
                    <div class="feature-item">
                        <h3>Simple to Use</h3>
                        <p>Enjoy a fast, modern, and intuitive interface that makes data management a breeze.</p>
                    </div>
                </div>
                <div class="hero-cta">
                    <button id="cta-register" class="btn btn-primary">Get Started for Free</button>
                    <button id="cta-learn-more" class="btn btn-secondary">Learn More</button>
                </div>
            </div>
        </section>
    `;

    // Attach event listeners to buttons after rendering
    document.getElementById("cta-register").onclick = () => window.router.navigate("/register");
    document.getElementById("cta-learn-more").onclick = () => window.router.navigate("/about");
};

export default Home;