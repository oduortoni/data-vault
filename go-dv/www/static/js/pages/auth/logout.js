/*
 * file: go-dv/www/static/js/pages/auth/logout.js
 * description: Logs out the current user by invalidating cookies.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 */

const Logout = async () => {
    window.app.innerHTML = `<section><h2>Logging out...</h2><pre id="logout-output"></pre></section>`;
    const output = document.getElementById("logout-output");

    try {
        const res = await fetch("/auth/logout", {
            method: "POST",
            credentials: "include"
        });
        const data = await res.json();
        output.innerText = JSON.stringify(data, null, 2);
    } catch {
        output.innerText = "Failed to contact logout endpoint";
    }
};

export default Logout;
