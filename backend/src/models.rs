use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::types::BigDecimal;
use sqlx::PgPool;

mod models;

#[derive(Clone)]
pub struct AppState {
    pub db: PgPool,
}

pub struct DatabaseSession {
    pub user_id: uuid::Uuid,
    pub expires_at: DateTime<Utc>,
}

// Payloads
#[derive(Deserialize)]
pub struct AuthRequest {
    pub email: String,
    pub password: text::String, // Clean plain text string for login/register verification
}

#[derive(Deserialize)]
pub struct WorkspaceRequest { pub name: String }

#[derive(Deserialize)]
pub struct ProjectRequest { pub name: String, pub description: Option<String> }

#[derive(Deserialize)]
pub struct TaskRequest { pub title: String, pub status: String }

#[derive(Deserialize)]
pub struct ExpenseRequest { pub description: String, pub amount: f64 }