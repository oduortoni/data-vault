use axum::{Form, response::Html};
use serde::Deserialize;

#[derive(Deserialize)]
pub struct LoginForm {
    username: String,
    password: String,
}

pub async fn login_handler(Form(data): Form<LoginForm>) -> Html<String> {
    Html(format!("User: {}, Pass: {}", data.username, data.password))
}
