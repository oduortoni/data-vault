use axum::Form;
use axum::response::IntoResponse;
use serde::Deserialize;
use argon2::{Argon2, PasswordHasher};
use argon2::password_hash::{SaltString, rand_core::OsRng};
use axum::response::Html;
use std::fs;

pub async fn signup_view() -> Html<String> {
    let content = fs::read_to_string("templates/signup.html")
        .unwrap_or_else(|_| {
            "<h1>500 Internal Server Error</h1><p>Template not found</p>".to_string()
        });
    Html(content)
}


#[derive(Deserialize)]
pub struct SignupData {
    email: String,
    password: String,
}

pub async fn signup_handler(
    Form(data): Form<SignupData>,
) -> impl IntoResponse {

    let salt = SaltString::generate(&mut OsRng);
    let argon2 = Argon2::default();
    let password_hash = argon2
        .hash_password(data.password.as_bytes(), &salt)
        .unwrap()
        .to_string();

    // TODO: insert into DB (next step)
    println!("Create user {}: {}", data.email, password_hash);

    "Signup successful"
}
