document.addEventListener("DOMContentLoaded", () => {
    const root = document.getElementById("root");
    if (!root) return;
  
    root.innerHTML = `
      <h2>Auth SPA</h2>
      <div id="auth-box">
        <form id="login-form">
          <input type="email" id="login-email" placeholder="Email" required />
          <input type="password" id="login-password" placeholder="Password" required />
          <button type="submit">Login</button>
        </form>
  
        <form id="register-form">
          <input type="email" id="register-email" placeholder="Email" required />
          <input type="password" id="register-password" placeholder="Password" required />
          <button type="submit">Register</button>
        </form>
  
        <button id="logout-btn">Logout</button>
      </div>
  
      <pre id="response-output"></pre>
    `;
  
    const output = document.getElementById("response-output");
  
    const sendJSON = async (url, payload) => {
        console.log(url, payload)
      try {
        const res = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          credentials: "include",
          body: JSON.stringify(payload)
        });
  
        const data = await res.json();
        output.innerText = JSON.stringify(data, null, 2);
      } catch (err) {
        output.innerText = "Failed to connect to server";
      }
    };
  
    document.getElementById("login-form").addEventListener("submit", (e) => {
      e.preventDefault();
      sendJSON("/auth/login", {
        email: document.getElementById("login-email").value,
        password: document.getElementById("login-password").value
      });
    });
  
    document.getElementById("register-form").addEventListener("submit", (e) => {
      e.preventDefault();
      sendJSON("/auth/register", {
        email: document.getElementById("register-email").value,
        password: document.getElementById("register-password").value
      });
    });
  
    document.getElementById("logout-btn").addEventListener("click", async () => {
      const res = await fetch("/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      const data = await res.json();
      output.innerText = JSON.stringify(data, null, 2);
    });
  });
  