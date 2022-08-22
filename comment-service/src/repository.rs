#![allow(proc_macro_derive_resolution_fallback)]

use diesel;
use diesel::prelude::*;

use crate::models::Comment;
use crate::models::NewComment;
use crate::models::Rating;
use crate::models::NewRating;

use crate::schema::comments;
use crate::schema::comments::dsl::*;
use crate::schema::ratings;
use crate::schema::ratings::dsl::*;

pub fn get_comment(comment_id: i32, connection: &PgConnection) -> QueryResult<Comment> {
    comments::table.find(comment_id)
        .first(&*connection)
}

pub fn create_comment(new_comment: NewComment, conn: &PgConnection) -> QueryResult<Comment> {
    diesel::insert_into(comments::table)
        .values(&new_comment)
        .get_result(conn)
}

pub fn show_all_comments_for_video(video_id_provided: i32, connection: &PgConnection) -> QueryResult<Vec<Comment>>  {
    comments.filter(video_id.eq(video_id_provided))
        .load::<Comment>(&*connection)
}

pub fn show_all_reported_comments(connection: &PgConnection) -> QueryResult<Vec<Comment>>  {
    comments.filter(reported.eq(true))
        .load::<Comment>(&*connection)
}

pub fn delete_comment(comment_id: i32, connection: &PgConnection) -> QueryResult<usize> {
    diesel::delete(comments::table.find(comment_id))
        .execute(connection)
}

pub fn report_comment(comment_id: i32, connection: &PgConnection) -> QueryResult<Comment> {
    diesel::update(comments::table.find(comment_id))
        .set(reported.eq(true))
        .get_result(connection)
}

pub fn create_or_update_rating(new_rating: NewRating, conn: &PgConnection) -> QueryResult<Rating> {
    match ratings.filter(rating_owner_email.eq(&new_rating.rating_owner_email))
        .filter(rating_video_id.eq(&new_rating.rating_video_id))
        .select(rating_id)
        .first::<i32>(&*conn) {
            Ok(found_rating_id) => {
                diesel::update(ratings::table.find(found_rating_id))
                .set(rating.eq(new_rating.rating))
                .get_result(conn)
            },
            Err(_) => {
                println!("i get here!");
                diesel::insert_into(ratings::table)
                .values(&new_rating)
                .get_result(conn)
            }
        }
}

pub fn get_rating_for_video(video_id_provided: i32, connection: &PgConnection) -> QueryResult<f32> {
    match ratings.filter(rating_video_id.eq(video_id_provided))
        .select(rating)
        .load::<i32>(&*connection) {
            Ok(ratings_vec) => {
                if ratings_vec.len() == 0 {
                    return Ok(0.0);
                }

                let mut rating_sum = 0;
                for a_rating in ratings_vec.iter() {
                    rating_sum += a_rating;
                }
                Ok(rating_sum as f32 / ratings_vec.len() as f32)
            },
            Err(err) => Err(err)
        }
}

pub fn get_rating_for_user(user_email_provided: String, video_id_provided: i32, connection: &PgConnection) -> QueryResult<i32> {
    match ratings.filter(rating_owner_email.eq(user_email_provided))
        .filter(rating_video_id.eq(&video_id_provided))
        .select(rating)
        .first::<i32>(&*connection) {
            Ok(found_rating) => Ok(found_rating),
            Err(err) => Err(err)
        }
}