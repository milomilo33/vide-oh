use std::env;

use diesel::result::Error;
use rocket::http::Status;
use rocket::response::status;
use rocket_contrib::json::Json;

use crate::auth::MyJWTClaims;
use crate::connection::DbConn;
use crate::models::Comment;
use crate::models::NewComment;
use crate::models::Rating;
use crate::models::NewRating;

use crate::repository;

#[get("/comments/<video_id>")]
pub fn show_all_comments_for_video(video_id: i32, connection: DbConn, my_claims: MyJWTClaims) -> Result<Json<Vec<Comment>>, Status> {
    repository::show_all_comments_for_video(video_id, &connection)
        .map(|comments| Json(comments))
        .map_err(|_| Status::NotFound)
}

#[post("/comments", format ="application/json", data = "<new_comment>")]
pub fn create_comment(new_comment: Json<NewComment>, connection: DbConn) ->  Result<Status, Status> {
    println!("here 0 {}",&new_comment.body);
    repository::create_comment(new_comment.into_inner(), &connection)
        .map(|_| Status::Ok)
        .map_err(|_| Status::BadRequest)
}

// #[get("/<id>")]
// pub fn get_post(id: i32, connection: DbConn) -> Result<Json<Post>, Status> {
//     sample::repository::get_post(id, &connection)
//         .map(|post| Json(post))
//         .map_err(|error| error_status(error))
// }

// #[put("/<id>", format = "application/json", data = "<post>")]
// pub fn update_post(id: i32, post: Json<Post>, connection: DbConn) -> Result<Json<Post>, Status> {
//     sample::repository::update_post(id, post.into_inner(), &connection)
//         .map(|post| Json(post))
//         .map_err(|error| error_status(error))
// }

// #[delete("/<id>")]
// pub fn delete_post(id: i32, connection: DbConn) -> Result<status::NoContent, Status> {
//     sample::repository::delete_post(id, &connection)
//         .map(|_| status::NoContent)
//         .map_err(|error| error_status(error))
// }

// fn error_status(error: Error) -> Status {
//     match error {
//         Error::NotFound => Status::NotFound,
//         _ => Status::InternalServerError
//     }
// }