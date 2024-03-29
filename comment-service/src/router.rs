use rocket;

use crate::connection;
use crate::handler;

pub fn create_routes() {
    rocket::ignite()
        .manage(connection::init_pool())
        .mount("/api/comments",
               routes![
                    handler::show_all_comments_for_video,
                    handler::create_comment,
                    handler::show_all_reported_comments,
                    handler::delete_comment,
                    handler::report_comment,
                    handler::create_or_update_rating,
                    handler::get_rating_for_video,
                    handler::get_rating_for_user
                    ],
        ).launch();
}