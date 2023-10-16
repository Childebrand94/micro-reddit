-- Drop Tables
DROP TABLE IF EXISTS post_votes CASCADE;
DROP TABLE IF EXISTS comment_vote CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    registered_at timestamp with time zone NOT NULL default NOW()
);

-- Posts Table
CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    author_id bigserial NOT NULL,
    url text NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    updated_at timestamp with time zone NOT NULL default NOW(),
    CONSTRAINT fk_author_id FOREIGN KEY(author_id) REFERENCES users(id)
);

-- Comments Table
CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    post_id bigserial NOT NULL,
    author_id bigserial NOT NULL,
    parent_id bigserial,
    message text NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    CONSTRAINT fk_post_id FOREIGN KEY(post_id) REFERENCES posts(id),
    CONSTRAINT fk_author_id FOREIGN KEY(author_id) REFERENCES users(id),
    CONSTRAINT fk_parent_id FOREIGN KEY(parent_id) REFERENCES comments(id)
);

-- Post Votes Table
CREATE TABLE IF NOT EXISTS post_votes (
    id bigserial PRIMARY KEY,
    post_id bigserial NOT NULL,
    user_id bigserial NOT NULL,
    up_vote boolean NOT NULL,
    CONSTRAINT fk_post_id FOREIGN KEY(post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Comment Votes Table
CREATE TABLE IF NOT EXISTS comment_vote (
    id bigserial PRIMARY KEY,
    user_id bigserial NOT NULL,
    comment_id bigserial NOT NULL,
    up_vote boolean NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_comment_id FOREIGN KEY(comment_id) REFERENCES comments(id)
);
