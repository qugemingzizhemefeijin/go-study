drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
drop table if exists users;


CREATE TABLE users (
  id         INTEGER not null primary key AUTOINCREMENT,
  uuid       VARCHAR(64) not null unique,
  name       VARCHAR(255),
  email      VARCHAR(255) not null unique,
  password   VARCHAR(255) not null,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
);

CREATE TABLE sessions (
  id         INTEGER not null primary key AUTOINCREMENT,
  uuid       VARCHAR(64) not null unique,
  email      VARCHAR(255),
  user_id    INT not null,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
);

CREATE TABLE threads (
  id         INTEGER not null primary key AUTOINCREMENT,
  uuid       VARCHAR(64) not null unique,
  topic      TEXT,
  user_id    INT not null,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
);

CREATE TABLE posts (
  id         INTEGER not null primary key AUTOINCREMENT,
  uuid       VARCHAR(64) not null unique,
  body       TEXT,
  user_id    INT not null,
  thread_id  INT not null,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
);
