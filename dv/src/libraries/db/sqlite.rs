// db/mod.rs or db.rs
use sqlx::{sqlite::SqlitePoolOptions, SqlitePool};
use std::time::Duration;
use std::env;
use tokio::sync::OnceCell;
use uuid::Uuid;

#[derive(sqlx::FromRow)]
pub struct User {
    pub id: String,
    pub email: String,
    pub password_hash: String,
}

static DB_POOL: OnceCell<SqlitePool> = OnceCell::const_new();

// called once at startup to initialize the database pool
pub async fn sqlite_init() -> Result<(), sqlx::Error> {
    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = SqlitePoolOptions::new()
        .max_connections(5)
        .acquire_timeout(Duration::from_secs(5))
        .connect(&db_url)
        .await?;
    
    DB_POOL.set(pool).map_err(|_| sqlx::Error::Configuration("Database already initialized".into()))?;
    Ok(())
}

fn get_pool() -> &'static SqlitePool {
    DB_POOL.get().expect("Database not initialized. Call init_db() first.")
}

pub async fn create_user(email: String, password_hash: String) -> Result<String, sqlx::Error> {
    let user_id = Uuid::new_v4();
    let pool = get_pool();
    
    sqlx::query("INSERT INTO users (id, email, password_hash) VALUES (?, ?, ?)")
        .bind(user_id.to_string())
        .bind(email)
        .bind(password_hash)
        .execute(pool)
        .await?;
    
    Ok(user_id.to_string())
}

pub async fn get_user_by_email(email: &str) -> Result<Option<User>, sqlx::Error> {
    let pool = get_pool();
    
    let user = sqlx::query_as::<_, User>("SELECT * FROM users WHERE email = ?")
        .bind(email)
        .fetch_optional(pool)
        .await?;
    
    Ok(user)
}

pub async fn get_all_users() -> Result<Vec<User>, sqlx::Error> {
    let pool = get_pool();
    
    let users = sqlx::query_as::<_, User>("SELECT * FROM users")
        .fetch_all(pool)
        .await?;
    
    Ok(users)
}
