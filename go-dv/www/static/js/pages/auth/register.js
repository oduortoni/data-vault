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
            username: form["username"].value,
            email: form["email"].value,
            password: form["password"].value
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
                <div class="form-group">
                    <label for="username">Username</label>
                    <input type="text" id="username" name="username" placeholder="Email" required />
                </div>
                <div class="form-group">
                    <label for="email">Email</label>
                    <input type="email" id="email" name="email" placeholder="Email" required />
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" id="password" name="password" placeholder="Password" required />
                </div>
                <button type="submit">Register</button>
            </form>
            <pre id="register-output"></pre>
        </section>
    `;
};

export default Register;
