/*
 * file: go-dv/www/static/js/pages/auth/login.js
 * description: Renders the login form and handles login submission.
 * author: toni
 * date: 2025-06-28
 * version: 1.1.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const Login = async () => {
    window.app.innerHTML = LoginView();

    const output = document.getElementById("login-output");
    const form = document.getElementById("login-form");

    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        const payload = {
            email: form["login-email"].value,
            password: form["login-password"].value,
        };

        try {
            const res = await fetch("/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify(payload),
            });

            const data = await res.json();
            // output.innerText = JSON.stringify(data, null, 2);
            if (res.ok) {
                output.innerText = "Login successful! Redirecting...";
                // Redirect to dashboard or home page after successful login
                setTimeout(() => {
                    window.router.navigate("/dashboard");
                }, 1000);
            } else {
                output.innerText = `Error: ${data.message || "Login failed"}`;
            }
        } catch (err) {
            output.innerText = "Failed to connect to server";
            console.error(err);
        }
    });
};

const LoginView = () => {
    return `
        <section class="auth-form">
            <h2>Login</h2>
            <form id="login-form">
                <input type="email" id="login-email" placeholder="Email" required />
                <input type="password" id="login-password" placeholder="Password" required />
                <button type="submit">Login</button>
            </form>
            <pre id="login-output"></pre>
        </section>
    `;
};

export default Login;
