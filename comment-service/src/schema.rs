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
    ratings (id) {
        id -> Int4,
        owner_email -> Varchar,
        rating -> Int4,
        video_id -> Int4,
    }
}

allow_tables_to_appear_in_same_query!(
    comments,
    ratings,
);
