/*
 * file: go-dv/www/static/js/pages/auth/register.js
 * description: Renders the register form and handles user registration.
 * author: toni
 * date: 2025-06-28
 * version: 1.1.0
 * license: MIT
 */

const Register = async () => {
    window.app.innerHTML = RegisterView();

    const output = document.getElementById("register-output");
    const form = document.getElementById("register-form");

    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        const payload = {
            email: form["register-email"].value,
            password: form["register-password"].value
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
        } catch (err) {
            output.innerText = "Failed to connect to server";
            console.error(err);
        }
    });
};

const RegisterView = () => {
    return `
        <section class="auth-form">
            <h2>Register</h2>
            <form id="register-form">
                <input type="email" id="register-email" placeholder="Email" required />
                <input type="password" id="register-password" placeholder="Password" required />
                <button type="submit">Register</button>
            </form>
            <pre id="register-output"></pre>
        </section>
    `;
};

export default Register;
