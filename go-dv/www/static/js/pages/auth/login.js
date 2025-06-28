/*
 * file: go-dv/www/static/js/pages/auth/login.js
 * description: Renders the login form and handles login submission.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const Login = () => {
    window.app.innerHTML = `
        <section>
            <h2>Login</h2>
            <form id="login-form">
                <input type="email" id="login-email" placeholder="Email" required />
                <input type="password" id="login-password" placeholder="Password" required />
                <button type="submit">Login</button>
            </form>
            <pre id="login-output"></pre>
        </section>
    `;

    const output = document.getElementById("login-output");

    document.getElementById("login-form").addEventListener("submit", async (e) => {
        e.preventDefault();
        const payload = {
            email: document.getElementById("login-email").value,
            password: document.getElementById("login-password").value
        };

        try {
            const res = await fetch("/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify(payload)
            });
            const data = await res.json();
            output.innerText = JSON.stringify(data, null, 2);
        } catch {
            output.innerText = "Failed to connect to server";
        }
    });
};

export default Login;
