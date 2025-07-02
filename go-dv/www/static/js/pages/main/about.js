/*
 * file: go-dv/www/static/js/pages/main/about.js
 * description: This file is used to render the about view.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const About = () => {
    window.app.innerHTML = `
        <section>
            <h1>Data Vault – About</h1>

            <p>
                Data Vault is a secure, extensible web application designed to collect, manage,
                and restrict access to user-submitted data structures (like forms or tables).
                The backend is written in Go with optional Rust integration, and the frontend is
                a decoupled SPA that communicates via JSON APIs.
                <div>
                    <strong>Author:</strong> Toni<br>
                    <strong>Date:</strong> 2025-06-28<br>
                    <strong>Version:</strong> 1.0.0<br>
                    <strong>License:</strong> MIT<br>
                    <strong>Copyright:</strong> 2025 Toni<br>
                    <strong>Contact:</strong> oduortoni@gmail.com<br>
                </div>
            </p>


            <p>
                <ol>
                    <li>Requirements Analysis: <button id="requirements">Requirements</button></li>
                </ol>
            </p>
        </section>
    `;

    document.getElementById("requirements").onclick = () => window.router.navigate("/requirements");
};

export default About;