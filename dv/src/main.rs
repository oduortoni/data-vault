// main.rs
use dv::libraries::server::config::{ServerConfig};
use dv::libraries::server::server::{Server, start_server};
use axum::response::Html;

#[tokio::main]
async fn main() {
    dotenvy::dotenv().ok();

    let config = ServerConfig::new("10000");
    let mut server = Server::new();
    server.get("/", index);

    if let Err(e) = start_server(config, &server).await {
        eprintln!("Server error: {}", e);
        std::process::exit(1);
    }
}

async fn index() -> Html<&'static str> {
    Html(r#"<!DOCTYPE html>
    <html>
        <head>
            <title>Data Vault</title>
        </head>
        <body>
            <h1>Welcome to Data Vault</h1>
            <p>This is the root page of the Data Vault server.</p>
        </body>
    </html>"#)
}
