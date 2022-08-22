use rocket;

use crate::connection;
use crate::handler;

pub fn create_routes() {
    rocket::ignite()
        .manage(connection::init_pool())
        .mount("/api",
               routes![
                    handler::show_all_comments_for_video,
                    handler::create_comment,
                    ],
        ).launch();
}