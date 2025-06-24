use serde::Deserialize;
use argon2::{Argon2, PasswordHasher};
use argon2::password_hash::{SaltString, rand_core::OsRng};
use axum::response::Html;
use std::fs;
use axum::{
    extract::Form,
    response::{IntoResponse, Redirect},
    http::StatusCode,
};
use crate::libraries::db::sqlite as db;

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
    // Hash the password
    let salt = SaltString::generate(&mut OsRng);
    let argon2 = Argon2::default();
    let password_hash = argon2
        .hash_password(data.password.as_bytes(), &salt)
        .unwrap()
        .to_string();
    
    // Use the database module's function
    match db::create_user(data.email, password_hash).await {
        Ok(_user_id) => Redirect::to("/").into_response(),
        Err(e) => {
            eprintln!("Failed to create user: {}", e);
            (
                StatusCode::INTERNAL_SERVER_ERROR,
                "Signup failed",
            ).into_response()
        }
    }
}
