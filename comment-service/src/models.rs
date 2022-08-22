#![allow(proc_macro_derive_resolution_fallback)]

use crate::schema::comments;
use crate::schema::ratings;

use chrono::NaiveDateTime;

#[derive(Queryable, AsChangeset, Serialize, Deserialize, Debug, Insertable)]
#[table_name = "comments"]
pub struct Comment {
    pub id: i32,
    pub owner_email: String,
    pub body: String,
    pub reported: bool,
    pub video_id: i32,
    pub posted_at: NaiveDateTime,
}

#[derive(Insertable, Serialize, Deserialize)]
#[table_name="comments"]
pub struct NewComment {
    pub owner_email: String,
    pub body: String,
    pub video_id: i32,
}

#[derive(Queryable, AsChangeset, Serialize, Deserialize, Debug, Insertable)]
#[table_name = "ratings"]
pub struct Rating {
    pub rating_id: i32,
    pub rating_owner_email: String,
    pub rating: i32,
    pub rating_video_id: i32,
}

#[derive(Insertable, Serialize, Deserialize)]
#[table_name="ratings"]
pub struct NewRating {
    pub rating_owner_email: String,
    pub rating: i32,
    pub rating_video_id: i32,
}