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
    comments.filter(video_id.eq(video_id_provided))
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
