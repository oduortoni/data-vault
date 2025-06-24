// main.rs
use axum::{routing::get, Router};
use std::{env, net::SocketAddr};
use tower_http::services::ServeDir;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    // Load env vars from .env if present (for development)
    dotenvy::dotenv().ok();
    
    // Read port from env or default to 10000
    let port = env::var("PORT")
        .unwrap_or_else(|_| "10000".into())
        .parse::<u16>()
        .unwrap();
    
    let addr = SocketAddr::from(([0, 0, 0, 0], port));
    
    // Build routes
    let app = Router::new()
        .route("/", get(root_handler))
        .nest_service("/static", ServeDir::new("./static"));
    
    println!("🚀 Vaultform server running at http://{}/", addr);
    
    let listener = TcpListener::bind(&addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn root_handler() -> &'static str {
    "Vaultform backend is running."
}