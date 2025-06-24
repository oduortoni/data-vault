use axum::response::Html;

pub async fn index() -> Html<&'static str> {
    Html(r#"<!DOCTYPE html>
    <html>
        <head>
            <title>Data Vault</title>
            <link rel="stylesheet" href="/static/css/styles.css">
            <script src="/static/js/script.js"></script>
        </head>
        <body>
            <h1>Welcome to Data Vault</h1>
            <p>This is the root page of the Data Vault server.</p>
            <form action="/login" method="post">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
                <button type="submit">Login</button>
            </form>
        </body>
    </html>"#)
}
