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
            email: form["email"].value,
            password: form["password"].value,
        };

        try {
            const res = await fetch("/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify(payload),
            });

            const data = await res.json();
            if (res.ok) {
                output.innerText = "Login successful! Redirecting...";
                // Redirect to dashboard after successful login
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
        <div class="auth-container">
            <div class="form-card">
                <h1>Login</h1>
                <div id="login-output" class="login-output"></div>
                <form id="login-form" novalidate>
                    <div class="form-group">
                        <label for="email">Email</label>
                        <input type="email" id="email" name="email" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input type="password" id="password" name="password" required>
                    </div>
                    <button type="submit" class="btn btn-primary">Login</button>
                </form>
                <div class="form-footer">
                    <p>Don't have an account? <a href="/register" id="link-to-register" data-link>Sign Up</a></p>
                </div>
            </div>
        </div>
    `;
};

export default Login;
