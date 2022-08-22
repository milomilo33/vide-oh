table! {
    comments (id) {
        id -> Int4,
        owner_email -> Varchar,
        body -> Text,
        reported -> Bool,
        video_id -> Int4,
        posted_at -> Timestamp,
    }
}

table! {
    ratings (rating_id) {
        rating_id -> Int4,
        rating_owner_email -> Varchar,
        rating -> Int4,
        rating_video_id -> Int4,
    }
}

allow_tables_to_appear_in_same_query!(
    comments,
    ratings,
);
