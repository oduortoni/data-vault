/*
 * file: go-dv/www/static/js/pages/docs/requirements.js
 * description: Renders the page that displays the requirements analysis for the Data Vault project.
 * author: toni
 * date: 2025-07-02
 * version: 1.1.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const Requirements = () => {
    window.app.innerHTML = RequirementsView();
};

const RequirementsView = () => {
    return `
        <section>
            <h1>Data Vault – Requirements Analysis</h1>

            <p><strong>Author:</strong> Toni<br>
            <strong>Date:</strong> 2025-06-28<br>
            <strong>Version:</strong> 1.0.0<br>
            <strong>Contact:</strong> oduortoni@gmail.com</p>

            <h2>1. Project Overview</h2>
            <p>
                Data Vault is a secure, extensible web application designed to collect, manage,
                and restrict access to user-submitted data structures (like forms or tables).
                The backend is written in Go with optional Rust integration, and the frontend is
                a decoupled SPA that communicates via JSON APIs.
            </p>

            <h2>2. Goals & Objectives</h2>
            <table>
                <thead>
                <tr><th>Goal</th><th>Description</th></tr>
                </thead>
                <tbody>
                <tr><td>Modularity</td><td>Clean separation of concerns with domain, infrastructure, and UI layers.</td></tr>
                <tr><td>Security</td><td>JWT-based auth via HttpOnly cookies, XSS and CSRF prevention.</td></tr>
                <tr><td>Authentication</td><td>Email/password login, session refresh, and logout support.</td></tr>
                <tr><td>Dynamic Data</td><td>Allow users to define and submit structured data (form/table-like).</td></tr>
                <tr><td>Persistence</td><td>Pluggable storage with default GORM (SQLite).</td></tr>
                <tr><td>SPA-first</td><td>Client-side routing, API consumption, and dynamic rendering.</td></tr>
                </tbody>
            </table>

            <h2>3. Functional Requirements</h2>

            <h3>3.1 Authentication</h3>
            <table>
                <thead>
                <tr><th>Feature</th><th>Description</th></tr>
                </thead>
                <tbody>
                <tr><td>Register</td><td>Create a user using email and password.</td></tr>
                <tr><td>Login</td><td>Issue JWT tokens and set cookies.</td></tr>
                <tr><td>Logout</td><td>Clear session cookies and invalidate refresh tokens.</td></tr>
                <tr><td>Middleware</td><td>Protect routes and return JSON error responses.</td></tr>
                </tbody>
            </table>

            <h3>3.2 Users</h3>
            <table>
                <thead>
                <tr><th>Feature</th><th>Description</th></tr>
                </thead>
                <tbody>
                <tr><td>CRUD</td><td>Manage users via Create, Read, Update, Delete operations.</td></tr>
                <tr><td>Interface</td><td>Storage abstraction with GORM or in-memory implementations.</td></tr>
                <tr><td>Validation</td><td>Ensure valid email format and required fields.</td></tr>
                </tbody>
            </table>

            <h3>3.3 Data Vault / Forms</h3>
            <table>
                <thead>
                <tr><th>Feature</th><th>Description</th></tr>
                </thead>
                <tbody>
                <tr><td>Define Structure</td><td>Allow user-defined schema for dynamic data collection.</td></tr>
                <tr><td>Restricted Access</td><td>Role-based access control (RBAC) for data entries.</td></tr>
                <tr><td>Visualize</td><td>Support future data visualizations or summaries.</td></tr>
                </tbody>
            </table>

            <h2>4. Non-Functional Requirements</h2>
            <table>
                <thead>
                <tr><th>Category</th><th>Description</th></tr>
                </thead>
                <tbody>
                <tr><td>Security</td><td>HttpOnly cookies, CSRF protection, no localStorage.</td></tr>
                <tr><td>Scalability</td><td>Decoupled frontend and backend services.</td></tr>
                <tr><td>Portability</td><td>Runs via Docker with SQLite default backend.</td></tr>
                <tr><td>Testability</td><td>Unit testing of domain logic.</td></tr>
                <tr><td>Maintainability</td><td>Swappable layers and clean architecture.</td></tr>
                <tr><td>User Experience</td><td>SPA-first with responsive feedback and navigation.</td></tr>
                </tbody>
            </table>

            <h2>5. Architecture Summary</h2>
            <pre><code>.
            ├── go-dv/
            │   ├── cmd/             → Main app entry
            │   ├── internal/
            │   │   ├── auth/        → JWT, login, register, middleware
            │   │   ├── server/      → Custom router, handlers
            │   │   ├── users/       → DTOs and interfaces
            │   │   └── models/      → GORM/in-memory user storage
            │   └── www/
            │       └── static/      → JS SPA, CSS, assets
            ├── dv/                  → Optional Rust core
            ├── Makefile             → Build automation
            └── .gitignore
            </code></pre>

            <h2>6. Assumptions</h2>
            <ul>
                <li>The frontend is SPA-only, no server-side rendering.</li>
                <li>APIs return JSON-formatted error messages consistently.</li>
                <li>Email is unique and used to identify users.</li>
                <li>The user store can be swapped at runtime (mock/GORM).</li>
                <li>Admin or advanced access control is out of scope for v1.</li>
            </ul>

            <h2>7. Future Enhancements</h2>
            <ul>
                <li>Support for fully dynamic JSON schema per user/org.</li>
                <li>Data summaries and exports (CSV, PDF).</li>
                <li>Advanced RBAC or multi-tenant organization support.</li>
                <li>Multi-language support (i18n).</li>
                <li>Playwright or Cypress end-to-end testing setup.</li>
            </ul>
        </section>
    `;
};

export default Requirements;
