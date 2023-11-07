-- Drop Tables
DROP TABLE IF EXISTS post_votes CASCADE;
DROP TABLE IF EXISTS comment_votes CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;


-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password text NOT NULL, 
    registered_at timestamp with time zone NOT NULL default NOW()
);

-- Posts Table
CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    author_id bigint NOT NULL,
    title text NOT NULL,
    url text NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    updated_at timestamp with time zone NOT NULL default NOW(),
    CONSTRAINT fk_post_author_id FOREIGN KEY(author_id) REFERENCES users(id)
);

-- Comments Table
CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    post_id bigint NOT NULL,
    author_id bigint NOT NULL,
    parent_id bigint default NULL,
    message text NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    CONSTRAINT fk_comment_post_id FOREIGN KEY(post_id) REFERENCES posts(id),
    CONSTRAINT fk_comment_author_id FOREIGN KEY(author_id) REFERENCES users(id),
    CONSTRAINT fk_comment_parent_id FOREIGN KEY(parent_id) REFERENCES comments(id)
);

-- Post Votes Table
CREATE TABLE IF NOT EXISTS post_votes (
    id bigserial PRIMARY KEY,
    post_id bigint NOT NULL,
    user_id bigint NOT NULL,
    up_vote boolean NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    UNIQUE (post_id, user_id),
    CONSTRAINT fk_post_vote_post_id FOREIGN KEY(post_id) REFERENCES posts(id),
    CONSTRAINT fk_post_vote_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Comment Votes Table
CREATE TABLE IF NOT EXISTS comment_votes (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    comment_id bigint NOT NULL,
    up_vote boolean NOT NULL,
    created_at timestamp with time zone NOT NULL default NOW(),
    UNIQUE (comment_id, user_id),
    CONSTRAINT fk_comment_votes_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_comment_votes_comment_id FOREIGN KEY(comment_id) REFERENCES comments(id)
);

-- Session Table  
CREATE TABLE IF NOT EXISTS sessions (
    session_id UUID PRIMARY KEY,
    user_id bigint NOT NULL,
    CONSTRAINT fk_session_user_id FOREIGN KEY(user_id) REFERENCES users(id)
)
