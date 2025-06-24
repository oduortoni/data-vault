use std::net::SocketAddr;
use crate::libraries::utils::utils::{port};

pub struct ServerConfig {
    pub default_port: String,
}

impl ServerConfig {
    pub fn new(default_port: &str) -> Self {
        Self {
            default_port: default_port.to_string(),
        }
    }

    pub fn get_port(&self) -> u16 {
        port(&self.default_port)
    }

    pub fn get_address(&self) -> SocketAddr {
        let port = self.get_port();
        SocketAddr::from(([0, 0, 0, 0], port))
    }
}
