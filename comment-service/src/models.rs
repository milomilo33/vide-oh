use chrono::DateTime;

#[derive(Queryable, AsChangeset, Serialize, Deserialize, Debug, Insertable)]
#[table_name = "comments"]
pub struct Comment {
    pub id: i32,
    pub owner_email: String,
    pub body: String,
    pub reported: bool,
    pub video_id: i32,
    pub posted_at: DateTime,
}

#[derive(Insertable, Serialize, Deserialize)]
#[table_name="comments"]
pub struct NewComment {
    pub owner_email: String,
    pub body: String,
    pub reported: bool,
    pub video_id: i32,
    pub posted_at: DateTime,
}

#[derive(Queryable, AsChangeset, Serialize, Deserialize, Debug, Insertable)]
#[table_name = "ratings"]
pub struct Rating {
    pub id: i32,
    pub owner_email: String,
    pub rating: i32,
    pub video_id: i32,
}

#[derive(Insertable, Serialize, Deserialize)]
#[table_name="ratings"]
pub struct NewRating {
    pub owner_email: String,
    pub rating: i32,
    pub video_id: i32,
}