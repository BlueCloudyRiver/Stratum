use axum::{
    routing::{delete, get, post, put},
    Router,
};
use sqlx::PgPool;
use std::{env, net::TcpListener};
use tower_http::cors::{Any, CorsLayer};
use dotenv::dotenv;



#[tokio::main]
async fn main() {
    dotenv().ok();

    let db_url = env::var("DATABASE_URL")
        .expect("CRITICAL: Database url is required");

    let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());

    let pool = PgPool::connect(&db_url)
        .await
        .expect("Failed to connect to PostgreSQL Database");

    let state = models::AppState { db: pool };

    let cors = CorsLayer::new()
        .allow_origin(Any)
        .allow_methods(Any)
        .allow_headers(Any);

    let app = Router::new()
        // auth
        .route("/api/auth/register", post(handlers::auth::register_handler))
        .route("/api/auth/login", post(handlers::auth::login_handler))
        .route("/api/auth/logout", post(handlers::auth::logout_handler))

        // workspaces
        .route(
            "/api/workspaces",
            get(handlers::workspaces::get_workspaces)
                .post(handlers::workspaces::add_workspace),
        )
        .route(
            "/api/workspaces/:workspace_id",
            put(handlers::workspaces::update_workspace)
                .delete(handlers::workspaces::delete_workspace),
        )

        // projects
        .route(
            "/api/workspaces/:workspace_id/projects",
            get(handlers::projects::get_projects)
                .post(handlers::projects::add_project),
        )
        .route(
            "/api/workspaces/:workspace_id/project/:project_id",
            put(handlers::projects::update_project)
                .delete(handlers::projects::delete_project),
        )

        .layer(cors)
        .with_state(state);

    let addr = format!("0.0.0.0:{}", port);
    println!("Server starting on port: {}", port);

    let listener = TcpListener::bind(&addr).unwrap();

    axum::serve(listener, app.into_make_service())
        .await
        .unwrap();
}