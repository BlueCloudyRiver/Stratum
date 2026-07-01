use axum::{routing::get, Router};
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/hello", get(helloworld))
        .route("/users", get(users_get));

    let listener = TcpListener::bind("0.0.0.0:8080")
        .await
        .unwrap();

    axum::serve(listener, app)
        .await
        .unwrap();
}

async fn helloworld() -> &'static str {
    "Hello world"
}

async fn users_get()