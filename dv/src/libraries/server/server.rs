use axum::{Router, handler::Handler, routing::MethodRouter};
use tower_http::services::ServeDir;
use tokio::net::TcpListener;
use super::config::ServerConfig;

#[derive(Debug, Clone)]
pub struct Server {
    pub router: Router,
    pub routes: Vec<(String, String, MethodRouter)>,
}

impl Server {
    pub fn new() -> Self {
        Self {
            router: Router::new(),
            routes: Vec::new(),
        }
    }

    pub fn add_route<H, T>(&mut self, path: &str, method: &str, handler: H)
    where
        H: Handler<T, ()>,
        T: 'static,
    {
        let route = match method.to_uppercase().as_str() {
            "GET" => axum::routing::get(handler),
            "POST" => axum::routing::post(handler),
            "PUT" => axum::routing::put(handler),
            "DELETE" => axum::routing::delete(handler),
            "PATCH" => axum::routing::patch(handler),
            _ => {
                eprintln!("Unsupported HTTP method: {}", method);
                return;
            }
        };

        self.routes.push((
            path.to_string(),
            method.to_uppercase(),
            route,
        ));
    }

    // Convenience methods
    pub fn get<H, T>(&mut self, path: &str, handler: H) -> &mut Self
    where
        H: Handler<T, ()>,
        T: 'static,
    {
        self.add_route(path, "GET", handler);
        self
    }

    pub fn post<H, T>(&mut self, path: &str, handler: H) -> &mut Self
    where
        H: Handler<T, ()>,
        T: 'static,
    {
        self.add_route(path, "POST", handler);
        self
    }
}


pub fn create_app(server: &Server) -> Router {
    // begin with the server's router which is embellished with the database issues
    let mut router = server.router.clone();
    
    // add all routes from the server
    for (path, _method, route) in &server.routes {
        router = router.route(path, route.clone());
    }
    
    router
}

pub async fn start_server(config: ServerConfig, server: &Server) -> Result<(), Box<dyn std::error::Error>> {
    let addr = config.get_address();

    let app = create_app(&server)
        .nest_service("/static", ServeDir::new("./static"));

    println!("🚀 Vaultform server running at http://{}/", addr);

    let listener = TcpListener::bind(&addr).await?;
    axum::serve(listener, app).await?;
    Ok(())
}
