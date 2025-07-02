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

        const password = form["password"].value;
        const confirmPassword = form["confirm-password"].value;
        if (password !== confirmPassword) {
            output.innerText = "Passwords do not match";
            return;
        }

        const payload = {
            username: form["username"].value,
            email: form["email"].value,
            password: form["password"].value,
            confirm: form["confirm-password"].value,
        };

        console.log(payload);
        console.log(JSON.stringify(payload))

        try {
            const res = await fetch("/auth/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify(payload)
            });

            const data = await res.json();
            if (res.ok) {
                output.innerText = "Registration successful! Redirecting...";
                // Redirect to login after successful registration
                setTimeout(() => {
                    window.router.navigate("/login");
                }, 1000);
            } else {
                output.innerText = `Error: ${data.message || "Registration failed"}`;
            }

        } catch (err) {
            output.innerText = "Failed to connect to server";
            console.error(err);
        }
    });
};

const RegisterView = () => {
    return `
        <div class="auth-container">
            <div class="form-card">
                <h1>Create Account</h1>
                <div id="register-output" class="register-output"></div>
                <form id="register-form" novalidate>
                <div class="form-group">
                        <label for="username">Username</label>
                        <input type="text" id="username" name="username" required>
                    </div>
                    <div class="form-group">
                        <label for="email">Email</label>
                        <input type="email" id="email" name="email" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input type="password" id="password" name="password" minlength="8" required>
                    </div>
                    <div class="form-group">
                        <label for="confirm_password">Confirm Password</label>
                        <input type="password" id="confirm-password" name="confirm-password" required>
                    </div>
                    <button type="submit" class="btn btn-primary">Sign Up</button>
                </form>
                <div class="form-footer">
                    <p>Already have an account? <a href="/login" id="link-to-login" data-link>Login</a></p>
                </div>
            </div>
        </div>
    `;
};

export default Register;
