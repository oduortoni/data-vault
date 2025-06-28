/*
 * file: go-dv/www/static/js/pages/auth/logout.js
 * description: Logs out the current user by invalidating cookies.
 * author: toni
 * date: 2025-06-28
 * version: 1.1.0
 * license: MIT
 */

const Logout = async () => {
    window.app.innerHTML = LogoutView();

    const output = document.getElementById("logout-output");

    try {
        const res = await fetch("/auth/logout", {
            method: "POST",
            credentials: "include",
        });

        const data = await res.json();
        output.innerText = JSON.stringify(data, null, 2);
    } catch (err) {
        output.innerText = "Failed to contact logout endpoint";
        console.error(err);
    }
};

const LogoutView = () => {
    return `
        <section class="auth-logout">
            <h2>Logging out...</h2>
            <pre id="logout-output">Please wait...</pre>
        </section>
    `;
};

export default Logout;
