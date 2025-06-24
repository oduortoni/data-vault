use std::env;

/*
* port/1
* loads the port from the environment variable PORT. Uses def_port as a fallback
* args:
*   def_port
*
* rets:
*   u16 because port numbers are unsigned 16-bit integers
*/
pub fn port(def_port: &str) -> u16 {
    env::var("PORT")
        .unwrap_or_else(|_| def_port.to_string())
        .parse::<u16>()
        .unwrap_or_else(|_| panic!("Invalid port: {}", def_port))
}
