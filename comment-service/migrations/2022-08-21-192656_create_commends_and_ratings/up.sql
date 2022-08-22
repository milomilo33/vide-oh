-- Your SQL goes here
CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  owner_email VARCHAR NOT NULL,
  body TEXT NOT NULL,
  reported BOOLEAN NOT NULL DEFAULT 'f',
  video_id INTEGER NOT NULL,
  posted_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE ratings (
  rating_id SERIAL PRIMARY KEY,
  rating_owner_email VARCHAR NOT NULL,
  rating INTEGER NOT NULL,
  rating_video_id INTEGER NOT NULL,
  CHECK (rating BETWEEN 1 AND 5)
);