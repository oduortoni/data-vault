/*
 * file: go-dv/www/static/js/pages/auth/register.js
 * description: Renders the register form and handles user registration.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 */

const Register = () => {
    window.app.innerHTML = `
        <section>
            <h2>Register</h2>
            <form id="register-form">
                <input type="email" id="register-email" placeholder="Email" required />
                <input type="password" id="register-password" placeholder="Password" required />
                <button type="submit">Register</button>
            </form>
            <pre id="register-output"></pre>
        </section>
    `;

    const output = document.getElementById("register-output");

    document.getElementById("register-form").addEventListener("submit", async (e) => {
        e.preventDefault();
        const payload = {
            email: document.getElementById("register-email").value,
            password: document.getElementById("register-password").value
        };

        try {
            const res = await fetch("/auth/register", {
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

export default Register;
