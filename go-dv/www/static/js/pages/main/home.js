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
                        <div class="feature-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line><polyline points="10 9 9 9 8 9"></polyline></svg>
                        </div>
                        <div class="feature-text">
                            <h3>Build Custom Forms</h3>
                            <p>Design forms and tables to capture the exact data you need, without any technical skills.</p>
                        </div>
                    </div>
                    <div class="feature-item">
                        <div class="feature-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
                        </div>
                        <div class="feature-text">
                            <h3>Secure & Private</h3>
                            <p>Your data is protected with robust security. You decide who can see and manage your information.</p>
                        </div>
                    </div>
                    <div class="feature-item">
                        <div class="feature-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.99 10.01l-1-4A1 1 0 0 0 19 5H5a1 1 0 0 0-1 1.01l-1 4A1 1 0 0 0 4 11h16a1 1 0 0 0 1-1.01zM4 11v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-8"></path><path d="M12 11v8"></path><path d="M9 19h6"></path></svg>
                        </div>
                        <div class="feature-text">
                            <h3>Simple to Use</h3>
                            <p>Enjoy a fast, modern, and intuitive interface that makes data management a breeze.</p>
                        </div>
                    </div>
                </div>
                <div class="hero-cta">
                <button id="cta-register" class="btn btn-primary" data-link>Get Started</button>
                <button id="cta-learn-more" class="btn btn-secondary" data-link>Learn More</button>
                </div>
            </div>
        </section>
    `;

    document.getElementById("cta-register").onclick = () => window.router.navigate("/register");
    document.getElementById("cta-learn-more").onclick = () => window.router.navigate("/about");
};

export default Home;