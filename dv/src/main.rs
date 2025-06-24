// main.rs
use dv::libraries::server::config::ServerConfig;
use dv::libraries::server::server::{Server, start_server};
use dv::web::handlers::login::{login_view, login_handler};
use dv::web::handlers::signup::{signup_view, signup_handler};
use dv::web::handlers::index::index;
use dv::libraries::db::sqlite::sqlite_init as dbinit;

#[tokio::main]
async fn main()  -> Result<(), Box<dyn std::error::Error>> {
    dotenvy::dotenv().ok();
    dbinit().await?;

    let config = ServerConfig::new("10000");
    let mut server = Server::new();

    server.get("/", index);
    server.get("/login", login_view);
    server.post("/login", login_handler);
    server.get("/signup", signup_view);
    server.post("/signup", signup_handler);

    if let Err(e) = start_server(config, &server).await {
        eprintln!("Server error: {}", e);
        std::process::exit(1);
    }

    Ok(())
}
