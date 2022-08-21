#[macro_use] extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[launch]
fn rocket() -> _ {
    let rocket = rocket::build();
    
    rocket.mount("/", routes![index])
}