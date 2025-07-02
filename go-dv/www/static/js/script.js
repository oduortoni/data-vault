/*
 * file: go-dv/www/static/js/script.js
 * description: Bootstraps the SPA and sets up client-side routing.
 * author: toni
 * date: 2025-06-28
 * version: 1.1.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

import { Router } from "./lib/index.js";
import { Dialog } from "./components/index.js";
import {
    Home,
    About,
    Register,
    Login,
    Logout,
    Dashboard,
    Requirements
} from "./pages/index.js";

document.addEventListener("DOMContentLoaded", () => {
    const app = document.getElementById("app");

    // Expose globally for convenience
    window.app = app;
    window.views = {
        Dialog,
        Home,
        About,
        Requirements,
        Dashboard,
        Login,
        Register,
        Logout,
    };

    const router = new Router();

    // Core routes
    router.fallback(Home);
    router.register("/", Home);
    router.register("/about", About);
    router.register("/requirements", Requirements);
    
    router.register("/dashboard", Dashboard);

    // Auth routes
    router.register("/login", Login);
    router.register("/register", Register);
    router.register("/logout", Logout);

    window.router = router;

    /**
     * Handles back/forward navigation
     */
    window.onpopstate = () => {
        router.navigate(location.pathname);
    };

    // Navbar links
    document.getElementById("nav-logo").onclick = () => router.navigate("/");
    document.getElementById("nav-home").onclick = () => router.navigate("/");
    document.getElementById("nav-about").onclick = () => router.navigate("/about");
    document.getElementById("nav-dashboard").onclick = () => router.navigate("/dashboard");
    document.getElementById("nav-login").onclick = () => router.navigate("/login");
    document.getElementById("nav-register").onclick = () => router.navigate("/register");
    document.getElementById("nav-logout").onclick = () => router.navigate("/logout");

    // Initial route on page load
    router.navigate(location.pathname);
});
