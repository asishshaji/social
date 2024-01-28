DROP TYPE IF EXISTS user_status CASCADE;

CREATE TYPE user_status AS ENUM('enabled', 'disabled', 'blacklisted');

DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username varchar(40) UNIQUE,
  bio varchar(180),
  avatar TEXT,
  hashed_password varchar(120) NOT NULL,
  company varchar(30) NOT NULL,
  status user_status NOT NULL DEFAULT 'enabled',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE groups (
  id SERIAL PRIMARY KEY,
  group_name VARCHAR(40) UNIQUE,
  group_bio VARCHAR(180),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  creator_id bigint REFERENCES users(id)
);

CREATE TABLE memberships (
  id SERIAL PRIMARY_KEY,
  user_id bigint REFERENCES users(id),
  group_id bigint REFERENCES groups(id),
  joined_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
);

DROP INDEX IF EXISTS idx_user_company;

CREATE INDEX idx_user_company ON users(company);