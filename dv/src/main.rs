// main.rs
use dv::libraries::server::config::{ServerConfig};
use dv::libraries::server::server::{Server, start_server};
use dv::web::handlers::login::login_handler;
use dv::web::handlers::index::index;

#[tokio::main]
async fn main() {
    dotenvy::dotenv().ok();

    let config = ServerConfig::new("10000");
    let mut server = Server::new();
    server.get("/", index);
    server.post("/login", login_handler);

    if let Err(e) = start_server(config, &server).await {
        eprintln!("Server error: {}", e);
        std::process::exit(1);
    }
}
