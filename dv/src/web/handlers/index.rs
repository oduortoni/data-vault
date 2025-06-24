use axum::response::Html;
use std::fs;

pub async fn index() -> Html<String> {
    let content = fs::read_to_string("templates/index.html")
        .unwrap_or_else(|_| {
            "<h1>500 Internal Server Error</h1><p>Template not found</p>".to_string()
        });
    Html(content)
}
