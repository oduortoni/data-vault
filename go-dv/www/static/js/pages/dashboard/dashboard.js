/*
 * file: go-dv/www/static/js/pages/dashboard/dashboard.js
 * description: This file is used to render the home view.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const Dashboard = () => {
    fetch("/dashboard", {
        method: "GET",
        credentials: "include",
    })
    .then(res => {
        if (!res.ok) {
            window.router.navigate("/login");
            return Promise.reject("Failed to fetch dashboard data");
        }
        return  res.json()
    })
    .then(user => {
        if (!user) {
            window.router.navigate("/login");
            return Promise.reject("No user data found");
        }
        DashboardView(user);
        return Promise.resolve(user);
    })
    .catch(err => {
        window.router.navigate("/login");
        return Promise.reject("Error fetching dashboard data or user not logged in", err);
    });
};

const DashboardView = ({username, email}) => {
    window.app.innerHTML = `
        <section class="hero">
            <h2>Welcome, ${username}</h2>
            <p>You registered with this email: ${email}!</p>
            <p>Here you can manage your account, view statistics, and more.</p>
        </section>
    `;
};

export default Dashboard;