#![feature(decl_macro, proc_macro_hygiene)]
#[macro_use]
extern crate diesel;
extern crate dotenv;
extern crate r2d2;
extern crate r2d2_diesel;
#[macro_use]
extern crate rocket;
extern crate rocket_contrib;
#[macro_use]
extern crate serde_derive;

use dotenv::dotenv;

mod schema;
mod connection;
mod models;
mod repository;
mod handler;
mod router;
mod auth;

fn main() {
    dotenv().ok();
    router::create_routes();
}