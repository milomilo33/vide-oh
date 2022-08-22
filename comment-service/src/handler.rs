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
pub fn create_comment(new_comment: Json<NewComment>, connection: DbConn, my_claims: MyJWTClaims) ->  Result<Status, Status> {
    if !my_claims.email.eq(&new_comment.owner_email) {
        return Err(Status::Unauthorized);
    }

    repository::create_comment(new_comment.into_inner(), &connection)
        .map(|_| Status::Ok)
        .map_err(|_| Status::BadRequest)
}

#[get("/comments/reported")]
pub fn show_all_reported_comments(connection: DbConn, my_claims: MyJWTClaims) -> Result<Json<Vec<Comment>>, Status> {
    if !my_claims.role.eq(&String::from("Administrator")) {
        return Err(Status::Unauthorized);
    }

    repository::show_all_reported_comments(&connection)
        .map(|comments| Json(comments))
        .map_err(|_| Status::NotFound)
}

#[get("/comments/delete/<comment_id>")]
pub fn delete_comment(comment_id: i32, connection: DbConn, my_claims: MyJWTClaims) -> Result<Status, Status> {
    match repository::get_comment(comment_id, &connection) {
        Ok(comment) => {
            if my_claims.role.eq(&String::from("RegisteredUser")) && !my_claims.email.eq(&comment.owner_email) {
                return Err(Status::Unauthorized);
            }

            repository::delete_comment(comment_id, &connection)
                .map(|_| Status::Ok)
                .map_err(|_| Status::NotFound)
        },
        Err(_) => Err(Status::NotFound)
    }
}

#[get("/comments/report/<comment_id>")]
pub fn report_comment(comment_id: i32, connection: DbConn, my_claims: MyJWTClaims) -> Result<Status, Status> {
    repository::report_comment(comment_id, &connection)
        .map(|_| Status::Ok)
        .map_err(|_| Status::NotFound)
}

#[post("/ratings", format ="application/json", data = "<new_rating>")]
pub fn create_or_update_rating(new_rating: Json<NewRating>, connection: DbConn, my_claims: MyJWTClaims) ->  Result<Status, Status> {
    if !my_claims.email.eq(&new_rating.rating_owner_email) {
        return Err(Status::Unauthorized);
    }

    repository::create_or_update_rating(new_rating.into_inner(), &connection)
        .map(|_| Status::Ok)
        .map_err(|_| Status::BadRequest)
}

#[get("/ratings/total/<video_id>")]
pub fn get_rating_for_video(video_id: i32, connection: DbConn, my_claims: MyJWTClaims) -> Result<Json<f32>, Status> {
    repository::get_rating_for_video(video_id, &connection)
        .map(|rat| Json(rat))
        .map_err(|_| Status::NotFound)
}

#[get("/ratings/user/<owner_email>/<video_id>")]
pub fn get_rating_for_user(owner_email: String, video_id: i32, connection: DbConn, my_claims: MyJWTClaims) -> Result<Json<i32>, Status> {
    if !my_claims.email.eq(&owner_email) {
        return Err(Status::Unauthorized);
    }

    repository::get_rating_for_user(owner_email, video_id, &connection)
        .map(|rat| Json(rat))
        .map_err(|_| Status::NotFound)
}
