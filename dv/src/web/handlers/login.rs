use axum::{Form, response::Html};
use serde::Deserialize;
use std::fs;

pub async fn login_view() -> Html<String> {
    let content = fs::read_to_string("templates/login.html")
        .unwrap_or_else(|_| {
            "<h1>500 Internal Server Error</h1><p>Template not found</p>".to_string()
        });
    Html(content)
}

#[derive(Deserialize)]
pub struct LoginForm {
    username: String,
    password: String,
}

pub async fn login_handler(
    Form(data): Form<LoginForm>
) -> Html<String> {
    Html(format!("User: {}, Pass: {}", data.username, data.password))
}
